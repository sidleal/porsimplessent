package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Sentence struct {
	Producao  string
	Level     string
	Text      string
	Sentence  string
	Paragraph string
	RawText   string
}

type Align struct {
	Producao  string
	Level     string
	TextA     string
	SentenceA string
	TextB     string
	SentenceB string
}

func main() {

	log.Println("starting aligns...")

	rawSentences := readFile("../../porsimples/porsimples_sentences.tsv")
	lines := strings.Split(rawSentences, "\n")

	sentences := []Sentence{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		line = strings.Replace(line, "\r", "", -1)
		tokens := strings.Split(line, "\t")
		// log.Println(tokens[0], tokens[1], tokens[2], tokens[3], tokens[4], tokens[5])
		sentence := Sentence{tokens[0], tokens[1], tokens[2], tokens[3], tokens[4], tokens[5]}
		sentences = append(sentences, sentence)

	}

	rawAligns := readFile("../../porsimples/porsimples_aligns.tsv")
	lines = strings.Split(rawAligns, "\n")

	aligns := []Align{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		line = strings.Replace(line, "\r", "", -1)
		tokens := strings.Split(line, "\t")
		// log.Println(tokens[0], tokens[1], tokens[2], tokens[3], tokens[4], tokens[5])
		align := Align{tokens[0], tokens[1], tokens[2], tokens[3], tokens[4], tokens[5]}
		aligns = append(aligns, align)

	}

	generateConcatFiles("ORI", "ori_nat", sentences, aligns)
	generateConcatFiles("NAT", "nat_str", sentences, aligns)
	generateLengthFiles("ORI", "ori_nat", sentences, aligns)
	generateLengthFiles("NAT", "nat_str", sentences, aligns)
	generateAllSplitsFiles("ORI", "ori_nat", sentences, aligns)
	generateAllSplitsFiles("NAT", "nat_str", sentences, aligns)
	generateNoSplitsFiles("ORI", "ori_nat", sentences, aligns)
	generateNoSplitsFiles("NAT", "nat_str", sentences, aligns)
	generateFilesOriStr("pss2_align_length")
	generateFilesOriStr("pss1_align_all_splits")
	generateFilesOriStr("pss3_align_no_splits")
	generateTriplets()

}

//concat
func generateConcatFiles(level string, fileName string, sentences []Sentence, aligns []Align) {

	f1, err := os.OpenFile("../../pss/pss0_align_concat_"+fileName+".tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f1.Close()

	_, err = f1.WriteString("production_id\tlevel\tchanged\tsplit\tsentence_text_from\tsentence_text_to\n")
	if err != nil {
		log.Println("ERRO", err)
	}

	oriTotalSentencas := 0
	oriUnchanged := 0
	oriChanged := 0
	oriDivided := 0
	oriUndivided := 0

	count := 1
	for _, item := range sentences {
		if item.Level == level {
			// if item.Producao != "83" {
			// 	continue
			// }

			log.Println("-----------------------------------------------")

			log.Println(count, item.Level, item.Producao, item.Text, item.Sentence, item.RawText)
			pairs := getPairs(aligns, sentences, item.Sentence)
			divided := ""
			pairStr := ""
			if len(pairs) > 1 {
				divided = "S"
				oriDivided++
			} else {
				divided = "N"
				oriUndivided++
			}

			for _, pair := range pairs {
				pairStr += pair.RawText + " "
			}
			count++

			item.RawText = strings.TrimSpace(item.RawText)
			pairStr = strings.TrimSpace(pairStr)

			if pairStr == "" {
				continue
			}

			oriTotalSentencas++

			log.Println(count, item.Level, item.Producao, item.Text, item.Sentence)
			log.Println(item.RawText)
			log.Println(pairStr)

			changed := ""
			if item.RawText == pairStr {
				changed = "N"
				oriUnchanged++
			} else {
				changed = "S"
				oriChanged++
			}
			level := "ORI->NAT"
			if item.Level == "NAT" {
				level = "NAT->STR"
			}

			line := item.Producao + "\t" + level + "\t" + changed + "\t" + divided + "\t" + item.RawText + "\t" + pairStr
			_, err := f1.WriteString(line + "\n")
			if err != nil {
				log.Println("ERRO", err)
			}
		}
	}

	log.Println("------------------------------------")
	log.Println("Total sentenças : ", oriTotalSentencas)
	log.Println("Total sentenças Simplificadas: ", oriChanged)
	log.Println("Total sentenças NÃO Simplificadas: ", oriUnchanged)
	log.Println("Total sentenças Divididas: ", oriDivided)
	log.Println("Total sentenças NÃO Divididas: ", oriUndivided)

	log.Println("------------------------------------")

}

func getPairs(aligns []Align, sentences []Sentence, id string) []Sentence {
	ret := []Sentence{}
	for _, item := range aligns {

		if item.SentenceA == id {
			// log.Println(item.SentenceB)
			if ok, sentence := getSentence(sentences, item.SentenceB); ok {
				ret = append(ret, sentence)
			}
		}
	}
	return ret

}

func getSentence(sentences []Sentence, id string) (bool, Sentence) {
	for _, item := range sentences {
		// log.Println("[", item.Sentence, "]", "[", id, "]")
		if item.Sentence == id {
			return true, item
		}
	}
	return false, Sentence{}

}

func readFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	// r := charmap.ISO8859_1.NewDecoder().Reader(f)
	r := io.Reader(f)

	ret := ""

	buf := make([]byte, 32*1024)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			ret += string(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
	}

	return ret

}

//length
func generateLengthFiles(level string, fileName string, sentences []Sentence, aligns []Align) {

	f1, err := os.OpenFile("../../pss/pss2_align_length_"+fileName+".tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f1.Close()

	_, err = f1.WriteString("production_id\tlevel\tchanged\tsplit\tsentence_text_from\tsentence_text_to\n")
	if err != nil {
		log.Println("ERRO", err)
	}

	oriTotalSentencas := 0
	oriUnchanged := 0
	oriChanged := 0
	oriDivided := 0
	oriUndivided := 0

	doubles := 0

	count := 1
	for _, item := range sentences {
		if item.Level == level {
			log.Println("--------------------------------------------------------------------------------")

			// log.Println(count, item.Level, item.Producao, item.Text, item.Sentence, item.RawText)
			pairs := getPairs(aligns, sentences, item.Sentence)
			divided := ""
			if len(pairs) > 1 {
				divided = "S"
				oriDivided++
			} else {
				divided = "N"
				oriUndivided++
			}

			pairStrs := getBestPair(item.RawText, pairs)
			if len(pairStrs) > 1 {
				doubles += (len(pairStrs) - 1)
			}

			for _, pairStr := range pairStrs {

				if pairStr == "" {
					continue
				}

				count++

				item.RawText = strings.TrimSpace(item.RawText)
				pairStr = strings.TrimSpace(pairStr)

				oriTotalSentencas++

				log.Println(count, item.Level, item.Producao, item.Text, item.Sentence)
				log.Println(item.RawText)
				log.Println(pairStr)

				changed := ""
				if item.RawText == pairStr {
					changed = "N"
					oriUnchanged++
				} else {
					changed = "S"
					oriChanged++
				}

				level := "ORI->NAT"
				if item.Level == "NAT" {
					level = "NAT->STR"
				}

				line := item.Producao + "\t" + level + "\t" + changed + "\t" + divided + "\t" + item.RawText + "\t" + pairStr
				_, err := f1.WriteString(line + "\n")
				if err != nil {
					log.Println("ERRO", err)
				}

				log.Println("-----")

				log.Println(line)
			}

		}
	}

	log.Println("------------------------------------")
	log.Println("Total PARES: ", oriTotalSentencas)
	log.Println("Sentenças repetidas (similaridades iguais): ", doubles)
	log.Println("------------------------------------")

}

type SentencePair struct {
	RawText    string
	Similarity int
}

type BySimilarity []SentencePair

func (a BySimilarity) Len() int           { return len(a) }
func (a BySimilarity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySimilarity) Less(i, j int) bool { return a[i].Similarity > a[j].Similarity }

func getBestPair(sentA string, sentBList []Sentence) []string {
	pairs := []SentencePair{}
	for _, item := range sentBList {
		log.Println("---", item.Sentence, item.RawText)

		tokens := strings.Split(strings.TrimSpace(item.RawText), " ")
		tokensLen := len(tokens)
		commonTokens := getQtyCommonTokens(sentA, item.RawText)
		log.Println("------------------->", tokensLen, commonTokens, tokensLen+commonTokens)

		itemSimilarity := tokensLen + commonTokens
		pairs = append(pairs, SentencePair{item.RawText, itemSimilarity})
	}

	sort.Sort(BySimilarity(pairs))

	bestPairs := []string{}

	lastSim := 0
	for _, item := range pairs {
		if item.Similarity >= lastSim {
			bestPairs = append(bestPairs, item.RawText)
			lastSim = item.Similarity
		}
	}

	return bestPairs

}

func getQtyCommonTokens(a string, b string) int {
	tokensA := strings.Split(a, " ")
	tokensB := strings.Split(b, " ")
	ret := 0
	for _, tokenB := range tokensB {
		for _, tokenA := range tokensA {
			if tokenA == tokenB {
				ret++
				break
			}
		}
	}
	return ret
}

//all
func generateAllSplitsFiles(level string, fileName string, sentences []Sentence, aligns []Align) {

	f1, err := os.OpenFile("../../pss/pss1_align_all_splits_"+fileName+".tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f1.Close()

	_, err = f1.WriteString("production_id\tlevel\tchanged\tsplit\tsentence_text_from\tsentence_text_to\n")
	if err != nil {
		log.Println("ERRO", err)
	}

	oriTotalSentencas := 0
	oriUnchanged := 0
	oriChanged := 0
	oriDivided := 0
	oriUndivided := 0

	count := 1
	for _, item := range sentences {
		if item.Level == level {
			log.Println("--------------------------------------------------------------------------------")

			// log.Println(count, item.Level, item.Producao, item.Text, item.Sentence, item.RawText)
			pairs := getPairs(aligns, sentences, item.Sentence)
			divided := ""
			if len(pairs) > 1 {
				divided = "S"
				oriDivided++
			} else {
				divided = "N"
				oriUndivided++
			}

			pairStrs := getAllPairs(item.RawText, pairs)

			for _, pairStr := range pairStrs {

				if pairStr == "" {
					continue
				}

				count++

				item.RawText = strings.TrimSpace(item.RawText)
				pairStr = strings.TrimSpace(pairStr)

				oriTotalSentencas++

				log.Println(count, item.Level, item.Producao, item.Text, item.Sentence)
				log.Println(item.RawText)
				log.Println(pairStr)

				changed := ""
				if item.RawText == pairStr {
					changed = "N"
					oriUnchanged++
				} else {
					changed = "S"
					oriChanged++
				}

				level := "ORI->NAT"
				if item.Level == "NAT" {
					level = "NAT->STR"
				}

				line := item.Producao + "\t" + level + "\t" + changed + "\t" + divided + "\t" + item.RawText + "\t" + pairStr
				_, err := f1.WriteString(line + "\n")
				if err != nil {
					log.Println("ERRO", err)
				}

				log.Println("-----")

				log.Println(line)
			}

		}
	}

	log.Println("------------------------------------")
	log.Println("Total PARES: ", oriTotalSentencas)
	log.Println("------------------------------------")

}

func getAllPairs(sentA string, sentBList []Sentence) []string {
	pairs := []string{}
	for _, item := range sentBList {
		log.Println("---", item.Sentence, item.RawText)
		pairs = append(pairs, item.RawText)
	}
	return pairs

}

type Pair struct {
	Production string
	TextA      string
	TextB      string
	Divided    bool
}

//ori-str
func generateFilesOriStr(alignType string) {

	f4, err := os.OpenFile("../../pss/"+alignType+"_ori_str.tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f4.Close()

	_, err = f4.WriteString("production_id\tlevel\tchanged\tsplit\tsentence_text_from\tsentence_text_to\n")
	if err != nil {
		log.Println("ERRO", err)
	}

	sizeSentOri := readFile("../../pss/" + alignType + "_ori_nat.tsv")
	lines := strings.Split(sizeSentOri, "\n")

	oriNatPairs := []Pair{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		pair := Pair{}
		pair.Production = tokens[0]
		pair.Divided = tokens[3] == "S"
		pair.TextA = tokens[4]
		pair.TextB = tokens[5]
		oriNatPairs = append(oriNatPairs, pair)

	}

	sizeSentNat := readFile("../../pss/" + alignType + "_nat_str.tsv")
	lines = strings.Split(sizeSentNat, "\n")

	natStrPairs := []Pair{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		pair := Pair{}
		pair.Production = tokens[0]
		pair.Divided = tokens[3] == "S"
		pair.TextA = tokens[4]
		pair.TextB = tokens[5]
		natStrPairs = append(natStrPairs, pair)

	}

	for _, itemNat := range natStrPairs {
		for _, itemOri := range oriNatPairs {
			if itemNat.TextA == itemOri.TextB && itemNat.Production == itemOri.Production {

				oriNatChanged := itemOri.TextA != itemOri.TextB
				natStrChanged := itemNat.TextA != itemNat.TextB

				log.Println("----------------")
				log.Println(itemOri.TextA)
				log.Println(itemOri.TextB)
				log.Println(itemNat.TextB)

				if oriNatChanged && natStrChanged {

					line := itemOri.Production + "\t" + "ORI->STR\tS\t"
					if itemOri.Divided || itemNat.Divided {
						line += "S\t"
					} else {
						line += "N\t"
					}
					line += itemOri.TextA + "\t"
					line += itemNat.TextB + "\n"

					_, err := f4.WriteString(line)
					if err != nil {
						log.Println("ERRO", err)
					}

				}

			}
		}
	}

}

//all
func generateNoSplitsFiles(level string, fileName string, sentences []Sentence, aligns []Align) {

	f1, err := os.OpenFile("../../pss/pss3_align_no_splits_"+fileName+".tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f1.Close()

	_, err = f1.WriteString("production_id\tlevel\tchanged\tsplit\tsentence_text_from\tsentence_text_to\n")
	if err != nil {
		log.Println("ERRO", err)
	}

	oriTotalSentencas := 0
	oriUnchanged := 0
	oriChanged := 0
	oriDivided := 0
	oriUndivided := 0

	count := 1
	for _, item := range sentences {
		if item.Level == level {
			log.Println("--------------------------------------------------------------------------------")

			// log.Println(count, item.Level, item.Producao, item.Text, item.Sentence, item.RawText)
			pairs := getPairs(aligns, sentences, item.Sentence)
			divided := ""
			if len(pairs) > 1 {
				divided = "S"
				oriDivided++
			} else {
				divided = "N"
				oriUndivided++
			}

			pairStrs := getAllPairs(item.RawText, pairs)

			for _, pairStr := range pairStrs {

				if pairStr == "" {
					continue
				}

				if divided == "S" { // ignore split
					break
				}

				count++

				item.RawText = strings.TrimSpace(item.RawText)
				pairStr = strings.TrimSpace(pairStr)

				oriTotalSentencas++

				log.Println(count, item.Level, item.Producao, item.Text, item.Sentence)
				log.Println(item.RawText)
				log.Println(pairStr)

				changed := ""
				if item.RawText == pairStr {
					changed = "N"
					oriUnchanged++
				} else {
					changed = "S"
					oriChanged++
				}

				level := "ORI->NAT"
				if item.Level == "NAT" {
					level = "NAT->STR"
				}

				line := item.Producao + "\t" + level + "\t" + changed + "\t" + divided + "\t" + item.RawText + "\t" + pairStr
				_, err := f1.WriteString(line + "\n")
				if err != nil {
					log.Println("ERRO", err)
				}

				log.Println("-----")

				log.Println(line)
			}

		}
	}

	log.Println("------------------------------------")
	log.Println("Total PARES: ", oriTotalSentencas)
	log.Println("------------------------------------")

}

func generateTriplets() {

	f4, err := os.OpenFile("../../pss/triplets_length.tsv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("ERRO", err)
	}
	defer f4.Close()

	sizeSentOri := readFile("../../pss/pss2_align_length_ori_nat.tsv")
	lines := strings.Split(sizeSentOri, "\n")

	oriNatPairs := []Pair{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		pair := Pair{}
		pair.Production = tokens[0]
		pair.TextA = tokens[4]
		pair.TextB = tokens[5]
		oriNatPairs = append(oriNatPairs, pair)

	}

	sizeSentNat := readFile("../../pss/pss2_align_length_nat_str.tsv")
	lines = strings.Split(sizeSentNat, "\n")

	natStrPairs := []Pair{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		pair := Pair{}
		pair.Production = tokens[0]
		pair.TextA = tokens[4]
		pair.TextB = tokens[5]
		natStrPairs = append(natStrPairs, pair)

	}

	_, err = f4.WriteString("production_id\tlevel\tchanged_ori_nat\tchanged_nat_str\toriginal_text\tnatural_text\tstrong_text\n")
	if err != nil {
		log.Println("ERRO", err)
	}

	for _, itemNat := range natStrPairs {
		for _, itemOri := range oriNatPairs {
			if itemNat.TextA == itemOri.TextB && itemNat.Production == itemOri.Production {

				oriNatChanged := itemOri.TextA != itemOri.TextB
				natStrChanged := itemNat.TextA != itemNat.TextB

				log.Println("----------------")
				log.Println(itemOri.TextA)
				log.Println(itemOri.TextB)
				log.Println(itemNat.TextB)

				line := itemOri.Production + "\t" + "ORI->NAT->STR\t"

				if oriNatChanged {
					line += "S\t"
				} else {
					line += "N\t"
				}

				if natStrChanged {
					line += "S\t"
				} else {
					line += "N\t"
				}

				line += itemOri.TextA + "\t"
				line += itemOri.TextB + "\t"
				line += itemNat.TextB + "\n"

				_, err := f4.WriteString(line)
				if err != nil {
					log.Println("ERRO", err)
				}

			}
		}
	}

}
