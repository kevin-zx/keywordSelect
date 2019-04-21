package main

import (
	"fmt"
	"github.com/kevin-zx/seotools/api_5118"
	"time"
)

//	FF9853E8CCB04E06902F54A047BBE453

func main() {
	rankMap := map[int]float64{
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
	domains := map[string]string{"智能家居": "www.cleveroom.com",
		"蒸汽石锅鱼加盟": "www.scmcld.com",
		"石锅鱼加盟连锁": "www.ynshiguoyu.com",
		"智能家居展":   "gz.smarthomeexpo.com.cn",
		"井用潜水泵":   "www.sileite.com",
		"水环式真空泵":  "www.julebengye.com",
		"排污泵":     "www.made-pump.com"}

	for k, d := range domains {
		keywordCount := 0
		pcPvSum := 0.00
		mobilePvSum := 0.00
		psResults, t, err := api_5118.ExportBaiduPcSearchResults(d, 1, "FF9853E8CCB04E06902F54A047BBE453")
		time.Sleep(time.Second * 2)
		if err != nil && err.Error() == "暂无数据" {
			continue
		}
		if err != nil {
			//panic(err)
			fmt.Println(err.Error())
			continue
		}
		if t == 1 {
			keywordCount = len(*psResults)
		} else {
			keywordCount = int(t) * 100
		}

		for _, pr := range *psResults {
			if pr.Rank != 0 {
				r := rankMap[pr.Rank%10]
				for i := 0; i < 4 && i < pr.Rank/10; i++ {
					r = r / 10.00
				}
				pcPvSum += float64(pr.BidwordPcpv) * r
				mobilePvSum += float64(pr.BidwordWisepv) * r
			}

		}
		fmt.Printf("%s,%s,%d,%f,%f\n", k, d, keywordCount, pcPvSum, mobilePvSum)
	}
}
