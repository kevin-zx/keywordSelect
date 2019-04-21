package main

import (
	"fmt"
	"github.com/kevin-zx/seotools/comm/baidu"
	"github.com/kevin-zx/seotools/comm/site_base"
	"jinzhunassist/domain"
	"net/url"
	"strings"
)

var keywordCount map[string]int

func main() {
	rootWords := []string{"苏州房产"}
	websites, err := GetMainWebByWords(rootWords)
	if err != nil {
		panic(err)
	}
	keywordCount = make(map[string]int)
	webKeywords := GetSiteKeywords(websites)
	webKeywords5118 := Get5118Keywords(websites)
	allKeywords := append(webKeywords, webKeywords5118...)
	countKeyword(allKeywords)
	for k, c := range keywordCount {
		fmt.Println(k, "----------", c)
	}

}

func GetMainWebByWords(rootWords []string) (webSites []string, err error) {
	var srs []baidu.SearchResult
	for _, rk := range rootWords {
		var rs *[]baidu.SearchResult
		rs, err = baidu.GetBaiduPcResultsByKeyword(rk, 1, 50)
		if err != nil {
			return
		}
		for _, r := range *rs {
			if r.IsHomePage() && r.RealUrl != "" && !strings.Contains(r.RealUrl, "baidu") {
				srs = append(srs, r)
			}
		}
	}

	for _, sr := range srs {
		webSites = append(webSites, sr.RealUrl)
	}

	return
}

func GetSiteKeywords(websites []string) (keywords []string) {
	for _, wurl := range websites {

		wi, err := site_base.ParseWebSeoFromUrl(wurl)
		if err != nil {
			continue
		}
		webKeywords := wi.SpiltKeywordsStr2Arr()
		if webKeywords == nil {
			continue
		}
		keywords = append(keywords, wi.SpiltKeywordsStr2Arr()...)

	}
	return
}

func Get5118Keywords(websites []string) (keywords []string) {
	for _, wurlStr := range websites {
		wurl, err := url.Parse(wurlStr)
		if err != nil {
			continue
		}
		siteDomain := wurl.Host
		si, err := domain.GetDomainInfo(siteDomain)
		if err != nil {
			continue
		}

		fmt.Println(si.MobilePvSum, si.KeywordCount)
		for _, keyword := range si.BaiduPCResult {
			keywords = append(keywords, keyword.Keyword)
		}
		totalPage := si.TotalPage
		for page := 2; page < totalPage && page <= 5; page++ {
			si, err = domain.GetDomainInfo(siteDomain)
			if err != nil {
				continue
			}

			fmt.Println(si.MobilePvSum, si.KeywordCount, "page2")
			for _, keyword := range si.BaiduPCResult {
				keywords = append(keywords, keyword.Keyword)
			}
		}

	}
	return
}

func countKeyword(ks []string) {
	for _, k := range ks {
		if k == "" {
			continue
		}
		if _, ok := keywordCount[k]; ok {
			keywordCount[k]++
		} else {
			keywordCount[k] = 1
		}
	}
}
