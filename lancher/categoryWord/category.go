package main

import (
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/baiduApiSDK/apiUtil"
	"github.com/kevin-zx/baiduApiSDK/baiduSDK"
	"github.com/kevin-zx/go-util/fileUtil"
	"github.com/kevin-zx/seotools/comm/baidu"
	"github.com/kevin-zx/seotools/comm/site_base"
	"github.com/kevin-zx/seotools/comm/urlhandler"
	"jinzhunassist/domain"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var keywordCount map[string]int
var csis map[string]CatSiteInfo
var keywordSearchCount map[string]int

type CatSiteInfo struct {
	SearchKeyword string
	SiteDomain    string
	SiteUrl       string
	SitePageInfo  site_base.WebPageSeoInfo
	SeoInfo5118   domain.SeoInfo
	Keywords      []string
}

func main() {
	siteDomain := "www.gametea.com"
	rootWords := []string{"麻将游戏", "游戏大厅", "棋牌游戏", "斗牛", "双扣", "拼十", "斗地主"}
	fileName := "data/" + siteDomain + ".csv"
	var rfile *os.File
	var err error
	if !fileUtil.CheckFileIsExist(fileName) {
		rfile, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}
	} else {
		rfile, err = os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	defer rfile.Close()
	csvWriter := csv.NewWriter(rfile)

	csis = make(map[string]CatSiteInfo)

	keywordSearchCount = make(map[string]int)
	websites, err := GetMainWebByWords(rootWords)
	if err != nil {
		panic(err)
	}
	var qrs []apiUtil.QueryResult
	for _, rs := range rootWords {
		fengChaoKeywords, err := GetBaiduFengchaoKeywords(rs)
		if err != nil {
			panic(err)
		}
		qrs = append(qrs, *fengChaoKeywords...)
	}
	keywordCount = make(map[string]int)
	webKeywords := GetSiteKeywords(websites)
	webKeywords5118 := Get5118Keywords(websites)

	allKeywords := append(webKeywords, webKeywords5118...)

	for _, fck := range qrs {
		allKeywords = append(allKeywords, fck.Word)
	}
	countKeyword(allKeywords)
	topKeywordsMap := SelectTopKeywords(50)
	var topKeywords []string
	for k := range topKeywordsMap {
		topKeywords = append(topKeywords, k)
	}
	run(topKeywords)

	for k, c := range keywordCount {
		err := csvWriter.Write([]string{k, strconv.Itoa(c)})
		if err != nil {
			panic(err)
		}
		fmt.Println(k, "----------", c)

	}
	csvWriter.Flush()

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

func GetBaiduFengchaoKeywords(word string) (keywords *[]apiUtil.QueryResult, err error) {
	ePAuthHeader := &baiduSDK.AuthHeader{
		Username: "baidu-酷讯2732150-7",
		Password: "Hotel^Kuxun789",
		Token:    "d0a3c5f9ea56ab0e4e73db39f9c8bc36",
		Action:   "API-SDK",
	}
	epandService := apiUtil.NewQueryExpandService(ePAuthHeader)
	keywords, err = epandService.ExpandWordsByQuery(word, 0)
	return
}

func Get5118Keywords(websites []string) (keywords []string) {
	for _, wurlStr := range websites {
		wurl, err := url.Parse(wurlStr)
		if err != nil {
			continue
		}
		siteDomain := wurl.Host
		si, err := domain.GetDomainInfo(siteDomain, 1)
		if err != nil {
			continue
		}

		fmt.Println(si.MobilePvSum, si.KeywordCount)
		for _, keyword := range si.BaiduPCResult {
			keywords = append(keywords, keyword.Keyword)
		}
		totalPage := si.TotalPage
		for page := 2; page < totalPage && page <= 5; page++ {
			si, err = domain.GetDomainInfo(siteDomain, page)
			if err != nil {
				continue
			}

			fmt.Println(si.MobilePvSum, si.KeywordCount, "page2")
			for _, keyword := range si.BaiduPCResult {
				keywords = append(keywords, keyword.Keyword)
			}
		}
		fmt.Println(wurlStr, ":", len(keywords))

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

func run(keywords []string) []string {
	var sKeywords []string
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

				si5118, err := domain.GetDomainInfo(sdomain, 1)
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
				sKeywords = append(sKeywords, csi.Keywords...)
				csis[sdomain] = csi
			}
		}

	}
	countKeyword(sKeywords)
	return sKeywords
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
