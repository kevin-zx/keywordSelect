package main

import (
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/go-util/fileUtil"
	"github.com/kevin-zx/seotools/seoinfoQuery"
	"os"
	"strconv"
)

func main() {
	tasks := [][]string{{"硅pu篮球场施工", "3jietiyu.com", "19"}, {"金刚网纱窗厂家", "www.sdjingshang.com", "3"}, {"博物馆展柜厂家", "www.zgtupian.com", "22"}, {"解剖模型", "www.cnbbmx.com", "2"}, {"英国海淘", "www.xclink.com", "3"}, {"北京场地布置", "www.bfhz5.com", "5"}, {"殡葬服务一条龙", "www.trbyfw.com", "-2"}, {"小型孵化机", "www.weihangfuhua.com", "39"}, {"电源适配器生产厂家", "www.soythink.com", "-2"}, {"济南叉车", "www.nljx.b2b.cn", "8"}, {"希腊移民", "www.timevisa.net", "30"}, {"金刚石磨块", "www.cghtmj.com", "2"}, {"Alevel培训", "www.xtgjedu.com", "7"}, {"停车场管理公司", "www.jxparking.com", "28"}, {"冷藏集装箱租赁", "www.nb-byd.com", "10"}, {"上海装潢公司", "www.sjqianxiang.com", "19"}, {"济南汽车租赁", "www.jinanjintai.com", "8"}, {"布袋脉冲除尘器", "www.bdcdccq.com", "4"}, {"短信开发包", "www.smsmodems.com", "5"}, {"二维动画制作", "www.ahlingzhi.com", "4"}, {"上海特种集装箱制造", "www.daqiaoth.com", "63"}, {"上海报废车回收", "www.shjdqcfw.com", "8"}, {"地板十大品牌", "www.homenice.com.cn", "23"}, {"智能停车公司", "www.xiziparking.com", "-2"}, {"钢琴出租", "www.pianorent.cn", "2"}, {"铝合金升降平台", "www.yzdsjj.com", "20"}, {"上海灌浆料", "www.shbzsy.com", "19"}, {"热固性粉末涂料", "www.wfqytl.com", "3"}, {"不锈钢盲板接头", "www.mshatuan.com", "39"}, {"土工模袋", "www.sdhrthb.com", "4"}, {"行星减速器", "www.xxjiansuji.net", "10"}, {"低温恒温槽", "www.shunliuyq.com", "3"}, {"隔音降噪", "www.99hbw.cn", "6"}, {"不发火混凝土", "www.tophaowan.com", "8"}, {"上海家庭装修哪家好", "www.haojiazs.com", "6"}, {"上海公积金提取", "www.gjj001.org.cn", "36"}, {"快速门厂家", "www.daoermen.com", "3"}, {"九江装修公司", "www.csrjhome.com", "23"}, {"智能家居", "www.cleveroom.com", "20"}, {"公园游乐设备", "www.yl198.cn", "6"}, {"西安月子会所", "www.xagzjd.com", "10"}, {"波纹管截止阀", "www.cndxv.com", "26"}, {"天然气lng供应", "www.lanranshunda.com", "66"}, {"拖链电缆", "www.yafeicable.cn", "11"}, {"办公室装修", "www.szjingshang.com", "121"}, {"铝压铸件", "www.hndxwj.com", "10"}, {"上海装修公司", "www.shuxin-sh.com", "-2"}, {"液晶拼接屏报价", "www.sz-landun.com", "6"}, {"企业信用管理", "www.hbjyp.com", "4"}, {"运动场地围网", "www.dicorlosports.com", "-2"}, {"河南不锈钢岗亭", "www.hnbya.com", "-2"}, {"齿轮油厂家", "www.sddl7.com", "48"}, {"北京阳光房", "www.lovelogson.com", "27"}, {"上海消防维保", "www.qgsy119.com", "-2"}, {"附近刮痧拔罐", "www.bjzhengyuantang.com", "52"}, {"郑州防水公司", "www.hnky888.com", "1"}, {"福耀汽车玻璃", "www.szfuyao.com", "8"}, {"澳洲房产墨尔本", "www.jalinrealty.com.cn", "24"}, {"砂磨机", "www.pingnuojx.com", "10"}, {"环氧富锌底漆", "www.jsblffcl.com", "23"}, {"北京活动执行公司", "www.cakmall.com", "2"}, {"北京画册设计公司", "www.pkuhe.cn", "4"}, {"郑州资质代办", "www.zizhidb.com", "12"}, {"彩色路面", "www.zbchangyuan.net", "4"}, {"润滑油生产厂家", "www.dochi.cn", "3"}, {"日本房产", "www.jpjuw.com", "7"}, {"无负压", "www.cszlgs.com", "4"}, {"透水混凝土", "www.belmay2006.com", "8"}, {"美式教育", "www.edub.us", "4"}, {"海淀网站建设", "www.hdwzjs.cn", "8"}, {"北京寿衣", "www.deshoufu.com", "15"}, {"拉伸气弹簧", "www.ydgasspring.com", "25"}, {"热泵厂家", "www.amitime.com.cn", "2"}, {"上海代理记账公司", "www.hqcw.com", "4"}, {"膜结构收费站", "www.sdyueda.com", "6"}, {"庭院景观设计公司", "www.shangtinggarden.com", "2"}, {"上海环保装修", "www.021sankui.com", "22"}, {"苏州软件公司", "www.dreamtek.net.cn", "3"}, {"短信猫", "www.wavecomcn.com", "14"}, {"蒸汽石锅鱼加盟", "www.scmcld.com", "4"}, {"青岛监控", "www.monit8532.com", "22"}, {"苏州离心机", "www.lixinji688.com", "-2"}, {"面粉石磨", "www.lylysmm.com", "55"}, {"超声波清洗器", "www.qtfsw.com", "-2"}, {"北京代理记账公司", "www.bjjdhr.com", "8"}, {"彩色沥青", "www.syjplm.com", "6"}, {"阳光玫瑰葡萄苗", "www.ygmgpt.com", "-2"}, {"北京活动策划公司", "www.baoatt.com", "4"}, {"丰鼎源消防玛钢沟槽管件", "www.bjfdyjc.com", "1"}, {"EPDM颗粒", "www.jtbrothersports.com", "32"}, {"保定保安公司", "www.bdjunxun.com", "4"}, {"河南建筑资质代办", "www.zzyzhgl.com", "2"}, {"会计公司", "www.bjjdlkj.com", "2"}, {"桃树苗价格", "www.tstsm.cn", "3"}, {"精品超市设计", "www.91jiyi.com", "8"}, {"防锈漆", "www.fangxiuqi365.com", "13"}, {"北京婚姻纠纷律师", "www.law-edu.com", "4"}, {"大变形金刚模型", "www.tieyunxing.cn", "6"}, {"办公室装修设计公司", "www.dw66.net", "14"}, {"上海快餐配送", "www.jiejiacy.com", "26"}}
	rcsvpath := "data/rc.csv"
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
	csvW := csv.NewWriter(rf)
	err = csvW.Write([]string{"keyword",
		"domain",
		"record",
		"record_home_page_index",
		"keyword_record",
		"keyword_record_home_page_index",
		"keyword_avg_title_match_rate",
		"keyword_avg_description_match_rate",
		"keyword_avg_title_full_match_count",
		"keyword_avg_description_full_match_count",
		"home_page_title_match_rate",
		"home_page_description_match_rate",
		"home_page_title_full_match_count",
		"home_page_description_full_match_count",
		"first_real_url",
		"first_title_match_rate",
		"first_description_match_rate",
		"first_title_full_match_count",
		"first_description_full_match_count",
		"home_page_title_keyword_match_rate",
		"home_page_description_keyword_match_rate",
		"first_title_keyword_match_rate",
		"first_description_keyword_match_rate",
		"avg_title_keyword_match_rate",
		"avg_description_keyword_match_rate",
		"rank"})
	if err != nil {
		panic(err)
	}
	for _, t := range tasks {
		taskRs, err := seoinfoQuery.MultiQuery([][]string{{t[1], t[0]}})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		taskR := taskRs[0]

		err = csvW.Write([]string{
			t[0],
			t[1],
			strconv.Itoa(taskR.SiteSeoInfo.Record),
			strconv.Itoa(taskR.SiteSeoInfo.RecordHomePageIndex),
			strconv.Itoa(taskR.KeywordSiteMatchInfo.KeywordRecord),
			strconv.Itoa(taskR.KeywordSiteMatchInfo.KeywordRecordHomePageIndex),
			fmt.Sprintf("%.3f", taskR.KeywordMatchInfo.KeywordAvgTitleLenPowerRate),
			fmt.Sprintf("%.3f", taskR.KeywordMatchInfo.KeywordAvgDescriptionLenPowerRate),
			fmt.Sprintf("%.3f", taskR.KeywordMatchInfo.KeywordAvgTitleFullMatchCount),
			fmt.Sprintf("%.3f", taskR.KeywordMatchInfo.KeywordAvgDescriptionFullMatchCount),
			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.HomePageMatchInfo.TitleMatchLenPowerRate),
			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.HomePageMatchInfo.DescriptionMatchLenPowerRate),
			strconv.Itoa(taskR.KeywordSiteMatchInfo.HomePageMatchInfo.TitleFullMatchCount),
			strconv.Itoa(taskR.KeywordSiteMatchInfo.HomePageMatchInfo.DescriptionFullMatchCount),
			taskR.KeywordSiteMatchInfo.FirstPageMatchInfo.RealUrl,
			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.FirstPageMatchInfo.TitleMatchLenPowerRate),
			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.FirstPageMatchInfo.DescriptionMatchLenPowerRate),
			strconv.Itoa(taskR.KeywordSiteMatchInfo.FirstPageMatchInfo.TitleFullMatchCount),
			strconv.Itoa(taskR.KeywordSiteMatchInfo.FirstPageMatchInfo.DescriptionFullMatchCount),
			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.HomePageMatchInfo.TitleKeywordMatchRate),
			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.HomePageMatchInfo.DescriptionKeywordMatchRate),

			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.FirstPageMatchInfo.TitleKeywordMatchRate),
			fmt.Sprintf("%.3f", taskR.KeywordSiteMatchInfo.FirstPageMatchInfo.DescriptionKeywordMatchRate),

			fmt.Sprintf("%.3f", taskR.KeywordMatchInfo.KeywordAvgTitleMatchRate),
			fmt.Sprintf("%.3f", taskR.KeywordMatchInfo.DescriptionAvgKeywordMatchRate),
			t[2],
		})
		if err != nil {
			panic(err)
		}
	}
	csvW.Flush()

}
