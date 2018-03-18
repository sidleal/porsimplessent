package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "railsanotador:pwd@(172.18.0.4)/anotador")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query(
		`select id, title from productions where project_id in (1,2) and status <> 'REMOVE' order by id `,
	)

	for rows.Next() {
		var id int64
		var title string
		err = rows.Scan(&id, &title)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}

		log.Println("---------------------------------")
		log.Println(id, title)
		treatProduction(db, id)

	}
	rows.Close()

	log.Println("done.")
}

func treatProduction(db *sql.DB, prodId int64) {
	rows, err := db.Query(
		fmt.Sprintf(`select id, tipo from textos where production_id = %d`, prodId),
	)
	mapSimp := make(map[string]int64)

	var id int64
	var tipo string
	for rows.Next() {
		err = rows.Scan(&id, &tipo)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		mapSimp[tipo] = id
	}

	log.Println(mapSimp)

	f1, err := os.OpenFile("../porsimples/porsimples_sentences.tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f1.Close()

	f2, err := os.OpenFile("../porsimples/porsimples_aligns.tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f2.Close()

	sentOri := getSentences(db, mapSimp["ORIGINAL"])
	for _, item := range sentOri {
		log.Println(prodId, "ORI", mapSimp["ORIGINAL"], item[0], item[1], item[2])
		_, err := f1.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\n", prodId, "ORI", mapSimp["ORIGINAL"], item[0], item[1], item[2]))
		if err != nil {
			log.Println("ERRO", err)
		}
	}

	sentNat := getSentences(db, mapSimp["NATURAL"])
	for _, item := range sentNat {
		log.Println(prodId, "NAT", mapSimp["NATURAL"], item[0], item[1], item[2])
		_, err := f1.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\n", prodId, "NAT", mapSimp["NATURAL"], item[0], item[1], item[2]))
		if err != nil {
			log.Println("ERRO", err)
		}
	}

	sentStr := getSentences(db, mapSimp["VIOLENTO"])
	for _, item := range sentStr {
		log.Println(prodId, "STR", mapSimp["VIOLENTO"], item[0], item[1], item[2])
		_, err := f1.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\n", prodId, "STR", mapSimp["VIOLENTO"], item[0], item[1], item[2]))
		if err != nil {
			log.Println("ERRO", err)
		}
	}

	alignsOri := getAligns(db, mapSimp["ORIGINAL"])
	for _, item := range alignsOri {
		log.Println(prodId, "ORI->NAT", mapSimp["ORIGINAL"], item[0], item[1], item[2])
		_, err := f2.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\n", prodId, "ORI->NAT", mapSimp["ORIGINAL"], item[0], item[1], item[2]))
		if err != nil {
			log.Println("ERRO", err)
		}
	}
	alignsNat := getAligns(db, mapSimp["NATURAL"])
	for _, item := range alignsNat {
		log.Println(prodId, "NAT->STR", mapSimp["NATURAL"], item[0], item[1], item[2])
		_, err := f2.WriteString(fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\n", prodId, "NAT->STR", mapSimp["NATURAL"], item[0], item[1], item[2]))
		if err != nil {
			log.Println("ERRO", err)
		}
	}

	rows.Close()
}

func getAligns(db *sql.DB, textId int64) [][]string {

	ret := [][]string{}

	rows, err := db.Query(
		fmt.Sprintf(`select sentenceA, textoB, sentenceB from alignments where texto_id = %d `, textId),
	)

	for rows.Next() {
		var sentA string
		var textB string
		var sentB string
		err = rows.Scan(&sentA, &textB, &sentB)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		ret = append(ret, []string{sentA, textB, sentB})
	}
	rows.Close()

	return ret
}

func getSentences(db *sql.DB, textId int64) [][]string {

	ret := [][]string{}

	rows, err := db.Query(
		fmt.Sprintf(`select id, paragraph from sentences where texto_id = %d `, textId),
	)

	for rows.Next() {
		var id string
		var par string
		err = rows.Scan(&id, &par)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}

		textSent := getSentence(db, id)

		ret = append(ret, []string{id, par, textSent})
	}
	rows.Close()

	return ret
}

func getSentence(db *sql.DB, id string) string {
	rows, err := db.Query(
		fmt.Sprintf(`select word from words where sentence_id = %v order by id`, id),
	)

	sent := ""
	for rows.Next() {
		var word string
		err = rows.Scan(&word)
		if err != nil {
			log.Println("Scan failed:", err.Error())
		}
		word = strings.TrimSpace(word)
		sent += strings.TrimSpace(word) + " "

	}
	rows.Close()

	sent = strings.Replace(sent, " $.", ".", -1)
	sent = strings.Replace(sent, " $,", ",", -1)
	sent = strings.Replace(sent, " $)", ")", -1)
	sent = strings.Replace(sent, " $]", "]", -1)
	sent = strings.Replace(sent, " $:", ":", -1)
	sent = strings.Replace(sent, " $;", ";", -1)
	sent = strings.Replace(sent, " $?", "?", -1)
	sent = strings.Replace(sent, " $!", "!", -1)
	sent = strings.Replace(sent, " $%", "%", -1)
	sent = strings.Replace(sent, "$\"", "\"", -1)
	sent = strings.Replace(sent, "$'", "'", -1)
	sent = strings.Replace(sent, "$( ", "(", -1)
	sent = strings.Replace(sent, "$[ ", "[", -1)
	sent = strings.Replace(sent, "$--", "--", -1)
	sent = strings.Replace(sent, "  ", " ", -1)

	regEx := regexp.MustCompile(`"\s(.*?)\s*?"`)
	sent = regEx.ReplaceAllString(sent, `"$1"`)

	return sent
}
