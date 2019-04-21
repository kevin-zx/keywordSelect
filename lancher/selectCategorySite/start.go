package main

import (
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/go-util/fileUtil"
	"github.com/kevin-zx/seotools/comm/baidu"
	"github.com/kevin-zx/seotools/comm/site_base"
	"github.com/kevin-zx/seotools/comm/urlhandler"
	"jinzhunassist/domain"
	"os"
	"strings"
)

var keywordSearchCount map[string]int
var csis map[string]CatSiteInfo
var keywordCount map[string]int
var csvW *csv.Writer

type CatSiteInfo struct {
	SearchKeyword string
	SiteDomain    string
	SiteUrl       string
	SitePageInfo  site_base.WebPageSeoInfo
	SeoInfo5118   domain.SeoInfo
	Keywords      []string
}

func main() {
	rcsvpath := "data/d6.csv"
	var rf *os.File
	var err error
	if !fileUtil.CheckFileIsExist(rcsvpath) {
		rf, err = os.Create(rcsvpath)
		if err != nil {
			panic(err)
		}
	} else {
		rf, err = os.OpenFile(rcsvpath, os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	defer rf.Close()
	csvW = csv.NewWriter(rf)
	keywordSearchCount = make(map[string]int)
	keywordCount = make(map[string]int)
	seedSite := "http://pazl.pingan.cn"
	SDomain, err := urlhandler.GetDomain(seedSite)
	if err != nil {
		panic(err)
	}
	si, err := site_base.ParseWebSeoFromUrl(seedSite)
	if err != nil {
		panic(err)
	}
	keywords := si.SpiltKeywordsStr2Arr()
	si5118, err := domain.GetDomainInfo(SDomain)
	if err != nil {
		panic(err)
	}
	for _, r5118 := range si5118.BaiduPCResult {
		keywords = append(keywords, r5118.Keyword)
	}
	keywords = RemoveDuplicatesAndEmpty(keywords)
	countKeyword(keywords)
	csis = make(map[string]CatSiteInfo)

	run(keywords)
	tkeyword := []string{}
	topkc := SelectTopKeywords(100)
	for k, c := range topkc {
		fmt.Println(k, c)
		tkeyword = append(tkeyword, k)
	}
	run(tkeyword)

	topkc = SelectTopKeywords(10000)
	for k, c := range topkc {
		fmt.Println(k, c)
	}

}

func SelectTopKeywords(i int) map[string]int {
	keywords := make(map[string]int)
	min := 1000000000
	for k, c := range keywordCount {
		if len(keywords) < i {
			keywords[k] = c
			if c < min {
				min = c
			}
		} else {
			if c > min {
				tmin := 10000000000
				m := false
				for kk, cc := range keywords {

					if cc == min && !m {
						delete(keywords, kk)
						keywords[k] = c
						if c < tmin {
							tmin = c
						}
						m = true

					} else if m {
						if cc == min {
							tmin = min
							break
						}
						if cc < tmin {
							tmin = cc
						}
					}
				}
				min = tmin
			}
		}

	}
	return keywords
}

func formatString(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "'", "", -1)
	return s
}

func SelectKeywords() []string {
	all := 0
	for _, c := range keywordCount {
		all += c
	}
	avg := float64(all) / float64(len(keywordCount))
	keywords := []string{}
	stand := float64(avg)
	for k, c := range keywordCount {
		if float64(c) > stand {
			keywords = append(keywords, k)
		}
	}
	return keywords
}

func run(keywords []string) []string {
	sKeywords := []string{}
	l := len(keywords)
	for i, k := range keywords {
		if _, ok := keywordSearchCount[k]; ok {
			continue
		}
		fmt.Printf("%d/%d\n", i, l)
		keywordSearchCount[k] = 1
		bmrs, err := baidu.GetBaiduPcResultsByKeyword(k, 1, 10)
		if err != nil {
			panic(err)
		}
		for _, bmr := range *bmrs {
			if bmr.RealUrl == "" {
				_ = bmr.GetPCRealUrl()
			}
			if bmr.RealUrl != "" && !strings.Contains(bmr.RealUrl, "baidu") {
				sdomain, err := urlhandler.GetDomain(bmr.RealUrl)
				if err != nil {
					fmt.Printf("%s-%s\n", bmr.RealUrl, err.Error())
					continue
				}
				if _, ok := csis[sdomain]; ok {
					continue
				}
				si, err := site_base.ParseWebSeoFromUrl(bmr.RealUrl)
				if err != nil {
					fmt.Printf("%s-%s\n", bmr.RealUrl, err.Error())
					continue
				}

				si5118, err := domain.GetDomainInfo(sdomain)
				if err != nil {
					fmt.Printf("%s-%s\n", bmr.RealUrl, err.Error())
					continue
				}
				csi := CatSiteInfo{SitePageInfo: *si, SearchKeyword: k, SiteUrl: bmr.RealUrl, SeoInfo5118: si5118, SiteDomain: sdomain}

				csi.Keywords = append(csi.SitePageInfo.SpiltKeywordsStr2Arr())
				//sKeywords = append(sKeywords,csi.SitePageInfo.SpiltKeywordsStr2Arr()...)
				for _, r58 := range csi.SeoInfo5118.BaiduPCResult {
					csi.Keywords = append(csi.Keywords, r58.Keyword)
				}
				cks := strings.Replace(strings.Join(csi.Keywords, "|"), ",", "", -1)
				err = csvW.Write([]string{csi.SiteDomain, csi.SiteUrl, csi.SearchKeyword, cks, formatString(csi.SitePageInfo.Title), formatString(csi.SitePageInfo.Description)})
				if err != nil {
					panic(err)
				}

				sKeywords = append(sKeywords, csi.Keywords...)
				csis[sdomain] = csi
			}
			csvW.Flush()
		}

	}
	countKeyword(sKeywords)
	return sKeywords
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

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	var keywordCount = make(map[string]int)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		duFlag := false
		for _, re := range ret {

			if len(a[i]) == 0 {
				duFlag = true
				break
			}
			if re == a[i] {
				if _, ok := keywordCount[re]; !ok {
					keywordCount[re] = 1
				}
				duFlag = true
				break
			}
		}
		if !duFlag {
			ret = append(ret, a[i])
		}
	}
	return
}
