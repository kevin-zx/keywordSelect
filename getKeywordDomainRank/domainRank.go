package getKeywordDomainRank

import (
	"github.com/kevin-zx/seotools/comm/baidu"
	"strings"
)

func GetPCRankDomain(keyword string, page int) (domains map[string]int, err error) {
	domains = make(map[string]int)
	brs, err := baidu.GetBaiduPcResultsByKeyword(keyword, page, 10)
	if err != nil {
		return
	}
	for _, br := range *brs {
		if br.DisplayUrl != "" && !strings.Contains(br.DisplayUrl, "baidu") && !strings.Contains(br.DisplayUrl, "...") {
			domain := br.DisplayUrl
			domain = strings.Replace(domain, "http://", "", -1)
			domain = strings.Replace(domain, "https://", "", -1)
			if strings.HasSuffix(domain, "/") {
				domain = string(domain[0 : len(domain)-1])
			}
			if strings.Contains(domain, "/") {
				continue
			}
			domains[domain+"_________________"+keyword] = br.Rank
			//domains = append(domains,domain)
		}

	}
	return
}

func GetMobileRankDomain(keyword string, page int) (domains map[string]int, err error) {
	domains = make(map[string]int)
	brs, err := baidu.GetBaiduMobileResultsByKeyword(keyword, page)
	if err != nil {
		return
	}
	for _, br := range *brs {
		if br.IsHomePage() {
			domain := strings.Replace(br.RealUrl, "http://", "", -1)
			domain = strings.Replace(domain, "https://", "", -1)
			domain = strings.Replace(domain, "/", "", -1)
			domains[domain+"_________________"+keyword] = br.Rank
		}

	}
	return
}
