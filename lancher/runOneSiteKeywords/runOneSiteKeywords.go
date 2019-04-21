package main

import (
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/seotools/seoinfoQuery"
	"os"
	"strings"
)

func main() {
	dnksf := "yuanhaow.csv"
	f, err := os.Open(dnksf)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	siteDomain := "yuanhaowang.com"
	allKeywords, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	task := [][]string{}
	for _, ak := range allKeywords {
		if len(ak) == 0 {
			continue
		}
		if len(ak[0]) == 0 {
			continue
		}
		task = append(task, []string{siteDomain, strings.Replace(ak[0], " ", "", -1)})
	}
	results, err := seoinfoQuery.MultiQuery(task, seoinfoQuery.PC)
	for _, r := range results {
		if r.KeywordSiteMatchInfo.KeywordRecord <= 1 || r.KeywordSiteMatchInfo.FirstPageMatchInfo.TitleKeywordMatchRate-r.KeywordMatchInfo.KeywordAvgTitleMatchRate <= -0.2 {
			fmt.Println(r.KeywordMatchInfo.Keyword)
		}
	}
}
