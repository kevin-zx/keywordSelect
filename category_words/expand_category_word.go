package category_words

import (
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/baiduApiSDK/apiUtil"
	"github.com/kevin-zx/baiduApiSDK/baiduSDK"
	"github.com/kevin-zx/go-util/fileUtil"
	"github.com/kevin-zx/keywordSelect/domain"
	"github.com/kevin-zx/seotools/comm/baidu"
	"github.com/kevin-zx/seotools/comm/site_base"
	"github.com/kevin-zx/seotools/comm/urlhandler"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetCategoryWords(siteDomain string, rootWords []string, apiKey5118 string) (fileName string, err error) {
	fileName = "./" + strings.Replace(siteDomain, ".", "_", -1) + "/cateWords.csv"
	path := strings.Replace(siteDomain, ".", "_", -1)
	rq, rfile, err := createFile(path, fileName)
	if err != nil {
		return
	}
	defer rfile.Close()
	if rq {
		return
	}

	csvWriter := csv.NewWriter(rfile)
	keywordCount := make(map[string]int)
	webSites, err := GetCategoryWebSite(rootWords)
	if err != nil {
		return
	}

	var allKeywords []string
	// 从凤巢拓词
	for _, rs := range rootWords {
		fengChaoKeywords, err := GetBaiduFengchaoKeywords(rs)
		if err != nil {
			panic(err)
		}
		for _, fck := range *fengChaoKeywords {
			allKeywords = append(allKeywords, fck.Word)
		}
	}
	// 从meta获取关键词
	webKeywords := GetSiteKeywords(webSites)
	// 从5118获取关键词
	webKeywords5118 := Get5118Keywords(webSites, apiKey5118)
	// 关键词集合起来
	allKeywords = append(allKeywords, webKeywords5118...)
	allKeywords = append(allKeywords, webKeywords...)
	countKeyword(allKeywords, &keywordCount)
	topKeywordsMap := SelectTopKeywords(50, &keywordCount)
	var topKeywords []string
	for k := range topKeywordsMap {
		topKeywords = append(topKeywords, k)
	}
	run(topKeywords, &keywordCount, 4, apiKey5118)
	for k, c := range keywordCount {
		err := csvWriter.Write([]string{k, strconv.Itoa(c)})
		if err != nil {
			panic(err)
		}
		fmt.Println(k, "----------", c)

	}
	csvWriter.Flush()
	return
}

func countKeyword(ks []string, keywordCount *map[string]int) {
	for _, k := range ks {
		if k == "" {
			continue
		}
		if _, ok := (*keywordCount)[k]; ok {
			(*keywordCount)[k]++
		} else {
			(*keywordCount)[k] = 1
		}
	}
}

func run(keywords []string, keywordCount *map[string]int, parallelismCount int, appkey string) {

	siteDomains := []string{}
	siteUrls := []string{}
	for i, k := range keywords {
		fmt.Println("keyword get:", i, "/", len(keywords))
		bmrs, err := baidu.GetBaiduPcResultsByKeyword(k, 1, 10)
		// 出错重查
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		for _, bmr := range *bmrs {
			if bmr.RealUrl == "" {
				_ = bmr.GetPCRealUrl()
			}
			if bmr.RealUrl != "" && !strings.Contains(bmr.RealUrl, "baidu") {
				sdomain, err := urlhandler.GetDomain(bmr.RealUrl)
				if err != nil || sdomain == "" {
					continue
				}
				siteDomains = append(siteDomains, sdomain)
				siteUrls = append(siteUrls, bmr.RealUrl)
			}
		}

	}

	taskChannel := make(chan string)
	resultChannel := make(chan []string)
	var sKeywords []string
	siteDomains = site_base.RemoveDuplicatesAndEmpty(siteDomains)
	siteUrls = site_base.RemoveDuplicatesAndEmpty(siteUrls)
	for i, sdomain := range siteDomains {
		fmt.Println("domain 5118 keywords get:", i, "/", len(siteDomains))
		si5118, err := domain.GetDomainInfo(sdomain, 1, appkey)
		if err != nil {
			fmt.Printf("%s-%s\n", sdomain, err.Error())
			continue
		}
		for _, k := range si5118.BaiduPCResult {
			sKeywords = append(sKeywords, k.Keyword)
		}

	}
	for i := parallelismCount; i > 0; i-- {
		go GetKeywordsFromWebUrl(taskChannel, resultChannel)
	}
	//siteUrls = site_base.RemoveDuplicatesAndEmpty(siteUrls)

	//发送任务
	go func() {
		for _, su := range siteUrls {
			taskChannel <- su
		}
	}()

	taskLen := len(siteUrls)
	for taskLen > 0 {
		select {
		case keywords := <-resultChannel:
			sKeywords = append(sKeywords, keywords...)
			taskLen--
			fmt.Println("siteUrl 还剩下:", taskLen)
		}
	}
	close(taskChannel)
	countKeyword(sKeywords, keywordCount)

}

func GetKeywordsFromWebUrl(taskChannel chan string, resultChannel chan []string) {
	for taskSiteUrl := range taskChannel {
		var result []string
		si, err := site_base.ParseWebSeoFromUrl(taskSiteUrl)
		if err == nil {
			result = si.SpiltKeywordsStr2Arr()
		}
		resultChannel <- result
	}
	//si, err := site_base.ParseWebSeoFromUrl(su)
	//if err != nil {
	//	fmt.Printf("%s-%s\n", su, err.Error())
	//	continue
	//}

}

func createFile(path string, fileName string) (recentQuery bool, rFile *os.File, err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		err = os.Mkdir(path, 0777) //0777也可以os.ModePerm
		if err != nil {
			return
		}
		err = os.Chmod(path, 0777)
		if err != nil {
			return
		}
	}

	if !fileUtil.CheckFileIsExist(fileName) {
		rFile, err = os.Create(fileName)
		if err != nil {
			return
		}
	} else {
		rFile, err = os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
		if err != nil {
			return
		}
		var finfo os.FileInfo
		finfo, err = rFile.Stat()
		if err != nil {
			return
		}
		hour := timeSub(finfo.ModTime(), time.Now())
		if hour < 10 {
			recentQuery = true
			fmt.Println("最近10个小时内已经查过了", fileName, "的行业词了")
			return
		}

	}
	return
}

func timeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 获取行业词
func GetCategoryWebSite(rootWords []string) (webSites []string, err error) {
	var srs []baidu.SearchResult
	for i, rk := range rootWords {
		fmt.Println("get category website", i+1, "/", len(rootWords))
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

func Get5118Keywords(websites []string, apiKey5118 string) (keywords []string) {
	for _, wurlStr := range websites {
		wurl, err := url.Parse(wurlStr)
		if err != nil {
			continue
		}
		siteDomain := wurl.Host
		si, err := domain.GetDomainInfo(siteDomain, 1, apiKey5118)
		if err != nil {
			continue
		}

		fmt.Println(si.MobilePvSum, si.KeywordCount)
		for _, keyword := range si.BaiduPCResult {
			keywords = append(keywords, keyword.Keyword)
		}
		totalPage := si.TotalPage
		for page := 2; page < totalPage && page <= 5; page++ {
			si, err = domain.GetDomainInfo(siteDomain, page, apiKey5118)
			if err != nil {
				continue
			}

			fmt.Println(si.MobilePvSum, si.KeywordCount, "page "+strconv.Itoa(page))
			for _, keyword := range si.BaiduPCResult {
				keywords = append(keywords, keyword.Keyword)
			}
		}
		fmt.Println(wurlStr, ":", len(keywords))

	}
	return
}

func SelectTopKeywords(i int, keywordCount *map[string]int) map[string]int {
	keywords := make(map[string]int)
	min := 1000000000
	for k, c := range *keywordCount {
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
