package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Sentence struct {
	Producao   string
	Level      string
	Text       string
	Splited    string
	Changed    string
	TextTarget string
}

type SentencePair struct {
	Producao string
	Level    string
	TextA    string
	TextB    string
	Splited  string
	Changed  string
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

//estatísticas
func main() {

	mapRepeat := map[string]int{}
	oriNatFile := readFile("../../pss/pss2_align_length_ori_nat.tsv")
	lines := strings.Split(oriNatFile, "\n")

	sentencesZH := 0
	sentencesFSP := 0

	oriSizeSentences := []Sentence{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		sentence := Sentence{}
		sentence.Producao = tokens[0]
		sentence.Level = tokens[1]
		sentence.Changed = tokens[2]
		sentence.Splited = tokens[3]
		sentence.Text = tokens[4]
		sentence.TextTarget = tokens[5]

		if _, ok := mapRepeat[sentence.Producao+"-"+sentence.Text]; !ok {
			oriSizeSentences = append(oriSizeSentences, sentence)
			mapRepeat[sentence.Producao+"-"+sentence.Text] = 1
			prod, _ := strconv.Atoi(sentence.Producao)
			if prod < 116 {
				sentencesZH++
			} else {
				sentencesFSP++
			}
		}

	}

	oriNatAllFile := readFile("../../pss/pss1_align_all_splits_ori_nat.tsv")
	lines = strings.Split(oriNatAllFile, "\n")

	oriNatAllSentences := []Sentence{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		sentence := Sentence{}
		sentence.Producao = tokens[0]
		sentence.Level = tokens[1]
		sentence.Changed = tokens[2]
		sentence.Splited = tokens[3]
		sentence.Text = tokens[4]
		sentence.TextTarget = tokens[5]

		oriNatAllSentences = append(oriNatAllSentences, sentence)

	}

	oriStrAllFile := readFile("../../pss/pss1_align_all_splits_ori_str.tsv")
	lines = strings.Split(oriStrAllFile, "\n")

	oriStrAllSentences := []Sentence{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		sentence := Sentence{}
		sentence.Producao = tokens[0]
		sentence.Level = tokens[1]
		sentence.Changed = tokens[2]
		sentence.Splited = tokens[3]
		sentence.Text = tokens[4]
		sentence.TextTarget = tokens[5]

		oriStrAllSentences = append(oriStrAllSentences, sentence)

	}

	mapRepeat = map[string]int{}
	natStrFile := readFile("../../pss/pss2_align_length_nat_str.tsv")
	lines = strings.Split(natStrFile, "\n")

	natSizeSentences := []Sentence{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		sentence := Sentence{}
		sentence.Producao = tokens[0]
		sentence.Level = tokens[1]
		sentence.Changed = tokens[2]
		sentence.Splited = tokens[3]
		sentence.Text = tokens[4]
		sentence.TextTarget = tokens[5]

		if _, ok := mapRepeat[sentence.Producao+"-"+sentence.Text]; !ok {
			natSizeSentences = append(natSizeSentences, sentence)
			mapRepeat[sentence.Producao+"-"+sentence.Text] = 1
		}

	}

	mapRepeat = map[string]int{}
	oriStrFile := readFile("../../pss/pss2_align_length_ori_str.tsv")
	lines = strings.Split(oriStrFile, "\n")

	numPairsOriStr := 0
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		if tokens[2] == "S" {
			numPairsOriStr++
		}
	}

	strSizeSentences := []Sentence{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		sentence := Sentence{}
		sentence.Producao = tokens[0]
		sentence.Level = tokens[1]
		sentence.Changed = tokens[2]
		sentence.Splited = tokens[3]
		sentence.Text = tokens[4]
		sentence.TextTarget = tokens[5]

		if _, ok := mapRepeat[sentence.Producao+"-"+sentence.Text]; !ok {
			strSizeSentences = append(strSizeSentences, sentence)
			mapRepeat[sentence.Producao+"-"+sentence.Text] = 1
		}

	}

	strFile := readFile("../../pss/pss1_align_all_splits_nat_str.tsv")
	lines = strings.Split(strFile, "\n")

	strSentences := []Sentence{}
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")
		sentence := Sentence{}
		sentence.Producao = tokens[0]
		sentence.Level = tokens[1]
		//sentence.Changed = tokens[2]
		//sentence.Splited = tokens[3]
		sentence.Text = tokens[5]

		strSentences = append(strSentences, sentence)

	}
	natStrAllSentences := strSentences

	triFile := readFile("../../pss/triplets_length.tsv")
	lines = strings.Split(triFile, "\n")

	triSameSentence := 0
	triSimplOriNat := 0
	triSimplNatStr := 0
	triSimplAll := 0
	triTotal := 0
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")

		if tokens[2] == "N" && tokens[3] == "N" {
			triSameSentence++
		}
		if tokens[2] == "S" && tokens[3] == "N" {
			triSimplOriNat++
		}
		if tokens[2] == "N" && tokens[3] == "S" {
			triSimplNatStr++
		}
		if tokens[2] == "S" && tokens[3] == "S" {
			triSimplAll++
		}
		triTotal++

	}

	allOriFile := readFile("../../pss/pss1_align_all_splits_ori_nat.tsv")
	lines = strings.Split(allOriFile, "\n")

	allNatFromSplit := 0
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")

		if tokens[3] == "S" {
			allNatFromSplit++
		}

	}

	allNatFile := readFile("../../pss/pss1_align_all_splits_nat_str.tsv")
	lines = strings.Split(allNatFile, "\n")

	allStrFromSplit := 0
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		tokens := strings.Split(line, "\t")

		if tokens[3] == "S" {
			allStrFromSplit++
		}

	}

	sameSentenceOriNat := 0
	for _, itemOri := range oriSizeSentences {
		for _, itemNat := range natSizeSentences {
			if itemOri.Producao == itemNat.Producao && itemOri.Text == itemNat.Text {
				if itemOri.Changed == "S" || itemOri.Splited == "S" {
					log.Println("ORI", itemOri.Producao, itemOri.Text)
				}
				sameSentenceOriNat++
			}
		}
	}

	sameSentenceNatStr := 0
	for _, itemNat := range natSizeSentences {
		for _, itemStr := range strSentences {
			if itemNat.Producao == itemStr.Producao && itemNat.Text == itemStr.Text {
				if itemNat.Changed == "S" || itemNat.Splited == "S" {
					log.Println("NAT", itemNat.Producao, itemNat.Text)
				}
				sameSentenceNatStr++
			}
		}
	}

	splitSentenceOriNat := 0
	for _, item := range oriSizeSentences {
		if item.Splited == "S" {
			splitSentenceOriNat++
		}
	}

	splitSentenceNatStr := 0
	for _, item := range natSizeSentences {
		if item.Splited == "S" {
			splitSentenceNatStr++
		}
	}

	simpNoSplitSentenceOriNat := 0
	for _, item := range oriSizeSentences {
		if item.Splited == "N" && item.Changed == "S" {
			simpNoSplitSentenceOriNat++
		}
	}

	simpNoSplitSentenceNatStr := 0
	for _, item := range natSizeSentences {
		if item.Splited == "N" && item.Changed == "S" {
			simpNoSplitSentenceNatStr++
		}
	}

	numPairsOriNat := 0
	for _, item := range oriSizeSentences {
		if item.Changed == "S" {
			numPairsOriNat++
		}
	}

	numPairsNatStr := 0
	for _, item := range natSizeSentences {
		if item.Changed == "S" {
			numPairsNatStr++
		}
	}

	mediaSimplificacoesSemDivisaoOri := 0
	minSimplificacoesSemDivisaoOri := 0
	maxSimplificacoesSemDivisaoOri := 0
	for _, item := range oriSizeSentences {
		if item.Splited == "N" && item.Changed == "S" {
			tokens := tokenizeText(item.TextTarget)
			mediaSimplificacoesSemDivisaoOri = (mediaSimplificacoesSemDivisaoOri + len(tokens)) / 2
			if len(tokens) > maxSimplificacoesSemDivisaoOri {
				maxSimplificacoesSemDivisaoOri = len(tokens)
			}
			if minSimplificacoesSemDivisaoOri == 0 || len(tokens) < minSimplificacoesSemDivisaoOri {
				minSimplificacoesSemDivisaoOri = len(tokens)
			}
		}
	}

	mediaSimplificacoesComDivisaoOri := 0
	minSimplificacoesComDivisaoOri := 0
	maxSimplificacoesComDivisaoOri := 0
	for _, item := range oriSizeSentences {
		if item.Splited == "S" {
			tokens := tokenizeText(item.TextTarget)
			mediaSimplificacoesComDivisaoOri = (mediaSimplificacoesComDivisaoOri + len(tokens)) / 2
			if len(tokens) > maxSimplificacoesComDivisaoOri {
				maxSimplificacoesComDivisaoOri = len(tokens)
			}
			if minSimplificacoesComDivisaoOri == 0 || len(tokens) < minSimplificacoesComDivisaoOri {
				minSimplificacoesComDivisaoOri = len(tokens)
			}
		}
	}

	mediaSimplificacoesSemDivisaoNat := 0
	minSimplificacoesSemDivisaoNat := 0
	maxSimplificacoesSemDivisaoNat := 0
	for _, item := range natSizeSentences {
		if item.Splited == "N" && item.Changed == "S" {
			tokens := tokenizeText(item.TextTarget)
			mediaSimplificacoesSemDivisaoNat = (mediaSimplificacoesSemDivisaoNat + len(tokens)) / 2
			if len(tokens) > maxSimplificacoesSemDivisaoNat {
				maxSimplificacoesSemDivisaoNat = len(tokens)
			}
			if minSimplificacoesSemDivisaoNat == 0 || len(tokens) < minSimplificacoesSemDivisaoNat {
				minSimplificacoesSemDivisaoNat = len(tokens)
			}
		}
	}

	mediaSimplificacoesComDivisaoNat := 0
	minSimplificacoesComDivisaoNat := 0
	maxSimplificacoesComDivisaoNat := 0
	for _, item := range natSizeSentences {
		if item.Splited == "S" {
			tokens := tokenizeText(item.TextTarget)
			mediaSimplificacoesComDivisaoNat = (mediaSimplificacoesComDivisaoNat + len(tokens)) / 2
			if len(tokens) > maxSimplificacoesComDivisaoNat {
				maxSimplificacoesComDivisaoNat = len(tokens)
			}
			if minSimplificacoesComDivisaoNat == 0 || len(tokens) < minSimplificacoesComDivisaoNat {
				minSimplificacoesComDivisaoNat = len(tokens)
			}
		}
	}

	mediaDiffSimplificacoesSemDivisaoOri := 0
	minDiffSimplificacoesSemDivisaoOri := 0
	maxDiffSimplificacoesSemDivisaoOri := 0
	for _, item := range oriSizeSentences {
		if item.Splited == "N" && item.Changed == "S" {
			tokensO := tokenizeText(item.Text)
			tokensT := tokenizeText(item.TextTarget)
			if len(tokensO)-len(tokensT) > 0 {
				mediaDiffSimplificacoesSemDivisaoOri = (mediaDiffSimplificacoesSemDivisaoOri + (len(tokensO) - len(tokensT))) / 2
				if (len(tokensO) - len(tokensT)) > maxDiffSimplificacoesSemDivisaoOri {
					maxDiffSimplificacoesSemDivisaoOri = (len(tokensO) - len(tokensT))
				}
				if minDiffSimplificacoesSemDivisaoOri == 0 || (len(tokensO)-len(tokensT)) < minDiffSimplificacoesSemDivisaoOri {
					minDiffSimplificacoesSemDivisaoOri = (len(tokensO) - len(tokensT))
				}
			}
		}
	}

	mediaDiffSimplificacoesComDivisaoOri := 0
	minDiffSimplificacoesComDivisaoOri := 0
	maxDiffSimplificacoesComDivisaoOri := 0
	for _, item := range oriSizeSentences {
		if item.Splited == "S" {
			tokensO := tokenizeText(item.Text)
			tokensT := tokenizeText(item.TextTarget)
			if len(tokensO)-len(tokensT) > 0 {
				mediaDiffSimplificacoesComDivisaoOri = (mediaDiffSimplificacoesComDivisaoOri + (len(tokensO) - len(tokensT))) / 2
				if (len(tokensO) - len(tokensT)) > maxDiffSimplificacoesComDivisaoOri {
					maxDiffSimplificacoesComDivisaoOri = (len(tokensO) - len(tokensT))
				}
				if minDiffSimplificacoesComDivisaoOri == 0 || (len(tokensO)-len(tokensT)) < minDiffSimplificacoesComDivisaoOri {
					minDiffSimplificacoesComDivisaoOri = (len(tokensO) - len(tokensT))
				}
			}
		}
	}

	pss1OriSize := 0
	for _, item := range oriNatAllSentences {
		if item.TextTarget != item.Text {
			pss1OriSize++
		}
	}
	pss1NatSize := 0
	for _, item := range natStrAllSentences {
		if item.TextTarget != item.Text {
			pss1NatSize++
		}
	}
	pss1StrSize := 0
	for _, item := range oriStrAllSentences {
		if item.TextTarget != item.Text {
			pss1StrSize++
		}
	}
	pss1Total := pss1OriSize + pss1NatSize + pss1StrSize

	pss2OriSize := 0
	for _, item := range oriSizeSentences {
		if item.TextTarget != item.Text {
			pss2OriSize++
		}
	}
	pss2NatSize := 0
	for _, item := range natSizeSentences {
		if item.TextTarget != item.Text {
			pss2NatSize++
		}
	}
	pss2StrSize := 0
	for _, item := range strSizeSentences {
		if item.TextTarget != item.Text {
			pss2StrSize++
		}
	}
	pss2Total := pss2OriSize + pss2NatSize + pss2StrSize

	pss3OriSize := 0
	for _, item := range oriSizeSentences {
		if item.TextTarget != item.Text && item.Splited == "N" {
			pss3OriSize++
		}
	}
	pss3NatSize := 0
	for _, item := range natSizeSentences {
		if item.TextTarget != item.Text && item.Splited == "N" {
			pss3NatSize++
		}
	}
	pss3StrSize := 0
	for _, item := range strSizeSentences {
		if item.TextTarget != item.Text && item.Splited == "N" {
			pss3StrSize++
		}
	}
	pss3Total := pss3OriSize + pss3NatSize + pss3StrSize

	log.Println("-------------------------")
	log.Println("Total sentences Original:", len(oriSizeSentences))
	log.Println("      Zero Hora:", sentencesZH)
	log.Println("      Caderno Ciencia FSP:", sentencesFSP)
	log.Println("Total sentences Natural:", len(natSizeSentences))
	log.Println("Total sentences Strong:", len(strSentences))
	log.Println("Total sentences ALL:", len(oriSizeSentences)+len(natSizeSentences)+len(strSentences))
	log.Println("")
	log.Println("Total sentences NO SIMPLIFICATION Original->Natural:", sameSentenceOriNat)
	log.Println("Total sentences NO SIMPLIFICATION Natural->Strong:", sameSentenceNatStr)
	log.Println("")
	log.Println("Total sentences SPLIT Original->Natural:", splitSentenceOriNat)
	log.Println("Total sentences SPLIT Natural->Strong:", splitSentenceNatStr)
	log.Println("")
	log.Println("Total sentences Natural from split:", allNatFromSplit)
	log.Println("Total sentences Strong from split:", allStrFromSplit)
	log.Println("")
	log.Println("Total sentences SIMPLIFIED (no split) Original->Natural:", simpNoSplitSentenceOriNat)
	log.Println("Total sentences SIMPLIFIED (no split) Natural->Strong:", simpNoSplitSentenceNatStr)
	log.Println("")
	log.Println("Total pairs simplified Original->Natural:", numPairsOriNat)
	log.Println("Total pairs simplified Natural->Strong:", numPairsNatStr)
	log.Println("Total pairs simplified Original->Strong:", numPairsOriStr)
	log.Println("Total all pairs simplified:", numPairsOriStr+numPairsNatStr+numPairsOriNat)
	log.Println("")
	log.Println("Total triplets NO SIMPLIFICATION 3 Levels:", triSameSentence)
	log.Println("Total triplets Simplified Only Original->Natural:", triSimplOriNat)
	log.Println("Total triplets Simplified Only Natural->Strong:", triSimplNatStr)
	log.Println("Total triplets Simplified 3 Levels:", triSimplAll)
	log.Println("Total triplets:", triTotal)
	log.Println("")
	log.Println("Mean token size of sentences - simplified (no split) - Ori->Nat:", mediaSimplificacoesSemDivisaoOri)
	log.Println("Min token size of sentences - simplified (no split) - Ori->Nat:", minSimplificacoesSemDivisaoOri)
	log.Println("Max token size tokens of sentences - simplified (no split) - Ori->Nat:", maxSimplificacoesSemDivisaoOri)
	log.Println("")
	log.Println("Mean token size of sentences - simplified (with split) - Ori->Nat:", mediaSimplificacoesComDivisaoOri)
	log.Println("Min token size of sentences - simplified (with split) - Ori->Nat:", minSimplificacoesComDivisaoOri)
	log.Println("Max token size tokens of sentences - simplified (with split) - Ori->Nat:", maxSimplificacoesComDivisaoOri)
	log.Println("")
	log.Println("Mean token size of sentences - simplified (no split) - Nat->Str:", mediaSimplificacoesSemDivisaoNat)
	log.Println("Min token size of sentences - simplified (no split) - Nat->Str:", minSimplificacoesSemDivisaoNat)
	log.Println("Max token size tokens of sentences - simplified (no split) - Nat->Str:", maxSimplificacoesSemDivisaoNat)
	log.Println("")
	log.Println("Mean token size of sentences - simplified (with split) - Nat->Str:", mediaSimplificacoesComDivisaoNat)
	log.Println("Min token size of sentences - simplified (with split) - Nat->Str:", minSimplificacoesComDivisaoNat)
	log.Println("Max token size tokens of sentences - simplified (with split) - Nat->Str:", maxSimplificacoesComDivisaoNat)
	log.Println("")
	log.Println("Mean tokens size diff of sentences - Originals vs simplified (no split) - Ori->Nat:", mediaDiffSimplificacoesSemDivisaoOri)
	log.Println("Min tokens size diff of sentences - Originals vs simplified (no split) - Ori->Nat:", minDiffSimplificacoesSemDivisaoOri)
	log.Println("Max tokens size diff of sentences - Originals vs simplified (no split) - Ori->Nat:", maxDiffSimplificacoesSemDivisaoOri)
	log.Println("")
	log.Println("Mean tokens size diff of sentences - Originals vs simplified (with split) - Ori->Nat:", mediaDiffSimplificacoesComDivisaoOri)
	log.Println("Min tokens size diff of sentences - Originals vs simplified (with split) - Ori->Nat:", minDiffSimplificacoesComDivisaoOri)
	log.Println("Max tokens size diff of sentences - Originals vs simplified (with split) - Ori->Nat:", maxDiffSimplificacoesComDivisaoOri)
	log.Println("")
	log.Println("Total PSS1 Original->Natural:", pss1OriSize)
	log.Println("Total PSS1 Natural->Strong:", pss1NatSize)
	log.Println("Total PSS1 Original->Strong:", pss1StrSize)
	log.Println("Total geral PSS1:", pss1Total)
	log.Println("")
	log.Println("Total PSS2 Original->Natural:", pss2OriSize)
	log.Println("Total PSS2 Natural->Strong:", pss2NatSize)
	log.Println("Total PSS2 Original->Strong:", pss2StrSize)
	log.Println("Total geral PSS2:", pss2Total)
	log.Println("")
	log.Println("Total PSS3 Original->Natural:", pss3OriSize)
	log.Println("Total PSS3 Natural->Strong:", pss3NatSize)
	log.Println("Total PSS3 Original->Strong:", pss3StrSize)
	log.Println("Total geral PSS3:", pss3Total)
	log.Println("")
	log.Println("-------------------------")

}

func tokenizeText(rawText string) []string {
	regEx := regexp.MustCompile(`([A-z]+)-([A-z]+)`)
	rawText = regEx.ReplaceAllString(rawText, "$1|hyp|$2")

	regEx = regexp.MustCompile(`\|gdot\|`)
	rawText = regEx.ReplaceAllString(rawText, ".")

	regEx = regexp.MustCompile(`\|gint\|`)
	rawText = regEx.ReplaceAllString(rawText, "?")

	regEx = regexp.MustCompile(`\|gexc\|`)
	rawText = regEx.ReplaceAllString(rawText, "!")

	regEx = regexp.MustCompile(`([\.\,"\(\)\[\]\{\}\?\!;:-]{1})`)
	rawText = regEx.ReplaceAllString(rawText, "  $1 ")

	regEx = regexp.MustCompile(`\s+`)
	rawText = regEx.ReplaceAllString(rawText, " ")

	return strings.Split(rawText, " ")
}
