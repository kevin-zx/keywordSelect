package domain

import (
	"github.com/kevin-zx/seotools/api_5118"
	"time"
)

var rankMap = map[int]float64{
	1: 0.8,
	2: 0.5,
	3: 0.4,
	4: 0.4,
	5: 0.3,
	6: 0.2,
	7: 0.1,
	8: 0.1,
	9: 0.1,
	0: 0.1,
}

type SeoInfo struct {
	KeywordCount  int
	PcPvSum       float64
	MobilePvSum   float64
	BaiduPCResult []api_5118.BaiduPCResult
	TotalPage     int
}

func GetDomainInfo(domain string, page int, appKey string) (si SeoInfo, err error) {
	si = SeoInfo{}
	time.Sleep(500 * time.Millisecond)
	psResults, t, err := api_5118.ExportBaiduPcSearchResults(domain, page, appKey)
	if err != nil && err.Error() == "暂无数据" {
		return
	}
	if err != nil {
		return
	}
	si.TotalPage = int(t)
	if t == 1 {
		si.KeywordCount = len(*psResults)
	} else {
		si.KeywordCount = int(t) * 100
	}

	for _, pr := range *psResults {
		if pr.Rank != 0 {
			r := rankMap[pr.Rank%10]
			for i := 0; i < 4 && i < pr.Rank/10; i++ {
				r = r / 10.00
			}
			si.PcPvSum += float64(pr.BidwordPcpv) * r
			si.MobilePvSum += float64(pr.BidwordWisepv) * r
		}

	}
	si.BaiduPCResult = *psResults
	return
}
