package category_words

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/kevin-zx/go-util/httpUtil"
	"github.com/kevin-zx/seotools/comm/site_base"

	//"github.com/kevin-zx/seotools/comm/urlhandler"
	"strings"
)

// 实现太难了，等到行业词完成再说
func GetCategoryWordsBySiteUrl(siteUrl string, domainKeywordApiKey5118 string, longwordApiKey5118 string) (fileName string, err error) {
	//domain,err := urlhandler.GetDomain(siteUrl)
	//if err != nil {
	//	panic(err)
	//}
	webCon, err := httpUtil.GetWebConFromUrl(siteUrl)
	if err != nil {
		webCon, err = httpUtil.GetWebConFromUrl(siteUrl)
	}
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(webCon))
	if err != nil {
		panic(err)
	}
	//doc.
	doc.Each(func(_ int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	text := doc.Text()
	fmt.Println(text)
	wi, _ := site_base.ParseWebSeoFromHtml(webCon)
	fmt.Println(wi)
	return
}
