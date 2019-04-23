package main

import (
	"encoding/csv"
	"fmt"
	"github.com/kevin-zx/go-util/fileUtil"
	"github.com/kevin-zx/seotools/seoinfoQuery"
	"jinzhunassist/domain"
	"jinzhunassist/getKeywordDomainRank"
	"os"
	"strconv"
	"strings"
)

func main() {

	rcsvpath := "data/rc_m3.csv"
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

	domains := []string{
		//"balidao.xutour.com","cd.liezx.com","daochengyading.wabuw.com","dl.liezx.com","domain","fz.liezx.com","geduan.miao-sou.com","gl.liezx.com","gy.liezx.com","hf.liezx.com","jiuzhaigou.xutour.com","jn.liezx.com","jx.cqbanjiafw.com","keyan.yizhishenbi.com","km.liezx.com","la.liezx.com","m.bwcx168.com","m.cdkspx.cn","m.cdszgc.cn","m.dafang24.com","m.jd1994.com.cn","m.jindianzics.com","m.jzskpw.com","m.kfsinfo.com","m.liangzhilaotao.com","m.maopaihuo.cn","m.qiannianshiguang.com","m.sbnzdm.com","m.scyongqin.com","m.whscits.com","m.zjskpw.com","m.zzdb88.com","maka.miao-sou.com",
		"maldives.xutour.com", "mauritius.xutour.com", "mc.cqczhj.com", "nc.liezx.com", "nj.liezx.com", "nn.liezx.com", "qd.liezx.com", "qj.liezx.com", "qz.liezx.com",
		"shanghai.gedu.org",
		"sjz.liezx.com", "srilanka.xutour.com", "sy.liezx.com", "sz.liezx.com", "ty.liezx.com", "tz.liezx.com", "wap.jician.com", "wap.zyx868.com", "wh.liezx.com", "www.001tee.com", "www.023yiliao.com", "www.027mjzsjt.com", "www.028bc.cn", "www.028dafa.com", "www.028jundu.com", "www.028twt.com", "www.1024sy.com", "www.15926296609.com", "www.21vis.com", "www.3nfs.cn", "www.519pr.com", "www.591jt.com", "www.59irv.com", "www.91yjn.com", "www.999pac.com", "www.aaaaa-kj.com", "www.adonschina.com", "www.afpfw.com", "www.ah-hxjy.com", "www.ahbkgd.cn", "www.ahxsl.com", "www.ahzyysh.com", "www.ajydpg.com", "www.akhj.com", "www.amysen.com.cn", "www.andalianbang.com", "www.ao-hua.com", "www.aohangsports.com", "www.arescloud.net", "www.backings.cn", "www.baiying808.com", "www.bandengnn.com", "www.banjiacq.net", "www.bebalance-c.com", "www.beiyuejx.com", "www.bestjie.cn", "www.bestomro.com", "www.bfirty.com", "www.bknyl.com", "www.bloomdata.cn", "www.bly1.com", "www.bokjqwzz.com", "www.bszc1.cn", "www.bwcx168.com", "www.byzcjy.com", "www.caitwlkj.com", "www.capakj.com", "www.cbledgs.com", "www.cddgms.com", "www.cdhao.cn", "www.cdhd168.com", "www.cdhdqd.cn", "www.cdhyjdsb.com", "www.cdkspx.cn", "www.cdmlsjr.com", "www.cdmolifang.com", "www.cdmshjy.com", "www.cdqt.cn", "www.cdsjgg.com", "www.cdszgc.cn", "www.cdtxft.com", "www.cdxczl.com", "www.cdxrnjd.com", "www.cdxycq.com", "www.cdzmjj.cn", "www.cfenglish.cn", "www.cfu101.com", "www.chbpc.cn", "www.chenghaomaoyi.com", "www.chenyalian.com", "www.chffm.cn", "www.china-chuban.com", "www.chinagbd.cn", "www.chinah1z1.com", "www.cklawyer.cn", "www.cleanout.com.cn", "www.clfcc.com", "www.clgsz.com", "www.cn-deman.com", "www.cq-bq.com", "www.cq5jdesign.com", "www.cqbanjiagongsi.com", "www.cqcq158.com", "www.cqduocai.com", "www.cqguoyi.cn", "www.cqhpcy.com", "www.cqhrjy.cn", "www.cqhtcw.cn", "www.cqhuayangyu.vip", "www.cqjdcs.cn", "www.cqjft.com", "www.cqjjyy.com", "www.cqjyxzs.com", "www.cqjzrg.cn", "www.cqlhmsm.com", "www.cqlmjpx.com", "www.cqlusi.com", "www.cqmjqt.com", "www.cqnfdr.com.cn", "www.cqnxqc.com", "www.cqnzj.com", "www.cqooooo.com", "www.cqqtjc.com", "www.cqsunrise.com", "www.cqtrzs.com", "www.cqtykj.cn", "www.cqukja.com", "www.cqxinyixin.com", "www.cqxzr.com", "www.cqyhgl.com", "www.cqyuelai.com.cn", "www.cqzkwx.com", "www.crppcb.cn", "www.cryun.com", "www.cuwedu.com", "www.cyopto.cn", "www.dabiaohuashi.com", "www.dafa028.com", "www.daishi.pro", "www.dajingui.com", "www.damonkids.cn", "www.daqipeixun.com", "www.delisibancai.com", "www.dhcy9.com", "www.dikazuche.com", "www.dmjgn.com", "www.dzstkjg.com", "www.euruni-sh.org", "www.face126.top", "www.fansicn.com", "www.feiy88.com", "www.fengfa56.com", "www.fengjr.com", "www.fgyaim.com", "www.fkxwz.com.cn", "www.forestbaby521.com", "www.fslgz.net", "www.gaokaozixun.cn", "www.good-zjj.com", "www.graspcm.com", "www.gsdb027.com", "www.gtinfrom.cn", "www.guangyitj.com", "www.gzhgbj.com", "www.gzmcsb.com", "www.hbbsajd.cn", "www.hbfdgd.com", "www.hbhycl.cn?", "www.hbjxsy.net", "www.hblxjzs.com", "www.hbracwyy.com", "www.hbxbyph.com", "www.hdtcpvcwd.com", "www.hefeizhihuixin.com", "www.helaser.com.cn", "www.helpsoft.com.cn", "www.hema.la", "www.henhenwanba.com", "www.heyuecar.com", "www.hfjingqiu.com", "www.hfzjx.com", "www.hms58.com", "www.hnjltpro.com", "www.hongweihj.com", "www.honmayi.cn", "www.htqlawyer.com", "www.huaaofs.com", "www.hybrothers.cn", "www.hzlab.net", "www.hzqsfz.com", "www.hzxdy.com", "www.ibzwh.com", "www.ichongdao.com", "www.iqushier.com", "www.itaojin.cn", "www.jbcjedu.com", "www.jiangudi.cn", "www.jiaodadianbo.com", "www.jiayibao.net", "www.jician.com", "www.jindianzics.com", "www.jinguanauto.com", "www.jinmaisoftware.com", "www.jinshizhongqing.com", "www.jljw.org", "www.jszp56.com", "www.jtcftd.com", "www.juntuotz.com", "www.jyfrt.com", "www.jyshie.com", "www.jzskpw.com", "www.k-kelite.com", "www.k5n.cn", "www.kalunguoji.com", "www.kangbange.com", "www.kardex.com.cn", "www.kmbot.com", "www.kmdwl.cn", "www.lggwy.com", "www.lianbai.com", "www.liangzhilaotao.com", "www.liequtuan.com", "www.liezx.com", "www.longwork.com.cn", "www.longyon.com", "www.lovejingling.com", "www.lpbdt.com", "www.lsbt888.cn", "www.lvfangtong9785.com", "www.lvrenwang.com", "www.lyglhg.com", "www.lywwyshp.com", "www.maopaihuo.cn", "www.mengzeyy.cn", "www.minitu.cn", "www.misseye.com.cn", "www.mjl198.com", "www.mjtsg.cn", "www.mlw113.com", "www.moguangfilm.com", "www.mrcen1.com", "www.mrzcw.com", "www.msdq027.com", "www.mshyk.cn", "www.mvrbak.com", "www.mwwe6t.com", "www.my-sj.cn", "www.n35vq1.com", "www.n3hoof.com", "www.n4qhvf.com", "www.n8g4up.com", "www.navnpt.com", "www.neto21.com", "www.newsound.cn", "www.nitrontech.cn", "www.nsemind.com", "www.outlook8.com", "www.owens.com.cn", "www.paiyin-print.com", "www.pengyuanda.net", "www.pgskpw.com", "www.poshsjd.com", "www.pudinuo-ceiling.com", "www.qbsmovie.com", "www.qchyzx.com", "www.qcpack.net", "www.qianyancanyin.com", "www.qijianzs.com", "www.qingmaicaiwu.cn", "www.qu-paper.com", "www.qzsxxx.com", "www.read-love.com", "www.rising-it.com", "www.robot-coat.com", "www.rxjyyjs.com", "www.sbnzdm.com", "www.scbidding.com", "www.sclinyou.com", "www.scrxjx.cn", "www.scxinhai.top", "www.scxsjs.cn", "www.sczbjz.cn", "www.sdspv.com", "www.sh-rongjing.com", "www.shan-pu.com", "www.shanghaijinlan.cn", "www.shangmenge.com", "www.shbaoyuan.com", "www.shhaokuo.com", "www.shjxzn.com", "www.shqmhb.cn", "www.shsaisong.com", "www.shsiom.com", "www.shtxwj.com", "www.shule9291.com", "www.shwht.com", "www.shyhjx021.com", "www.shyjxy.com", "www.sichuanks.cn", "www.sichuanks.com", "www.sichuanrr.com", "www.sinardhr.com", "www.sjlhcf.com", "www.sjolsd.com", "www.swucj.com", "www.sxarchery.com", "www.sy361.com", "www.szqzcpa.com", "www.szssj56.com", "www.szznzz.com", "www.tailingxidi.com", "www.tanyaxue.com", "www.teedq.com", "www.ticedu.cn", "www.tiemogs.com", "www.tiger-sh.com", "www.tioitio.com", "www.tjcaoshiyabo.com", "www.tongcaiedu.com", "www.tongjingwenquan.com", "www.tongou.net", "www.trumj.cn", "www.u-workshop.com", "www.u0vuch.com", "www.uni-technology.cn", "www.uphqp.com", "www.urducatena.com.cn", "www.wabuw.com", "www.wanweichengdu.com", "www.wanzhongheyi.com", "www.weikesong.com", "www.wenkaobang.com", "www.wenshang.net.cn", "www.wh-fxt.com", "www.wh-gsd.com", "www.wh-jy.com.cn", "www.wh898jnhb.cn", "www.whbcty.com", "www.whbfyf.com", "www.whbsgl.cn", "www.whbtkj.net", "www.whcyzx.cn", "www.whdell.cn", "www.whftpvc.com", "www.whgszz.cn", "www.whhtwc.com", "www.whjtty.cn", "www.whjtzscy.com", "www.whlnzs.cn", "www.whrjjt.cn", "www.whscits.com", "www.whxiyu.com", "www.whxtxbz.cn", "www.whzm.net", "www.whzybhs.com", "www.wmyzh.com", "www.wuhansujiefloor.com", "www.xgyndz.cn", "www.xianshengshe.com", "www.xiaochipeixun.com", "www.xljdwxb.com", "www.xnuo.com", "www.xsfmf.com", "www.xsw99.com", "www.xzlzj.net", "www.yameisj.com", "www.ybw666.com", "www.yifancw.com", "www.yiqunlc.com", "www.yizhishenbi.com", "www.ymxdl.cn", "www.yn3x.com", "www.ynhande.cn", "www.ynkingdee.com", "www.ynyonyou.com", "www.yongcanhuishou.com", "www.youxuanliuxue.com", "www.yqjshy.cn", "www.yuren2012.com", "www.yzhiyan.com", "www.zcjinchuang.com", "www.zcrpt.cn", "www.zeyhs.com", "www.zeyuedu.com", "www.zgcd8.com", "www.zgdoffice.com", "www.zgscqy.com", "www.zgsdcd.com", "www.zgzgzw.com", "www.zhaomiaopu.cn", "www.zhcwms.com", "www.zhengzhibangcd.com", "www.zheyi-art.com", "www.zhouheiya.net", "www.zixinkf.cn", "www.zjafdt.com", "www.zjskpw.com", "www.zmaaaaa.cn", "www.zzflk.com", "www2.whhzw.net", "wx.kmtqcw.cn", "xa.liezx.com", "yb.liezx.com",
	}
	var tasks [][]string
	var DRMap = make(map[string]int)
	for _, d := range domains {
		fmt.Println(d + "------------------------start")
		dis, err := domain.GetDomainInfo(d)
		if err != nil {
			//fmt.Printf("%s\n",err)
			continue
		}
		for i, di := range dis.BaiduPCResult {
			fmt.Println(i, len(dis.BaiduPCResult), "bpr")
			if di.BidwordWisepv >= 0 && di.BidwordWisepv < 200 {
				tMap, _ := getKeywordDomainRank.GetMobileRankDomain(di.Keyword, 3)

				t, _ := getKeywordDomainRank.GetMobileRankDomain(di.Keyword, 4)
				tMap = combineTwoMap(tMap, t)

				t, _ = getKeywordDomainRank.GetMobileRankDomain(di.Keyword, 5)
				tMap = combineTwoMap(tMap, t)
				t, _ = getKeywordDomainRank.GetMobileRankDomain(di.Keyword, 20)
				tMap = combineTwoMap(tMap, t)
				t, _ = getKeywordDomainRank.GetMobileRankDomain(di.Keyword, 30)
				tMap = combineTwoMap(tMap, t)
				t, _ = getKeywordDomainRank.GetMobileRankDomain(di.Keyword, 50)
				tMap = combineTwoMap(tMap, t)
				if err != nil {
					fmt.Printf("%s\n", err)
					continue
				}
				for d, _ := range tMap {
					if _, ok := DRMap[d]; ok {
						continue
					}
					dif, err := domain.GetDomainInfo(strings.Split(d, "_________________")[0])
					if err != nil {
						//fmt.Printf("%s\n",err)
						continue
					}

					if dif.MobilePvSum < 200 && dif.PcPvSum < 100 {
						tasks = append(tasks, []string{strings.Split(d, "_________________")[0], di.Keyword})
					}

				}
				DRMap = combineTwoMap(DRMap, tMap)
			}
		}
		mrs, err := seoinfoQuery.MultiQuery(tasks, seoinfoQuery.Mobile)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		for _, mr := range mrs {
			err := writeLine(*mr, *csvW, DRMap[mr.SiteSeoInfo.Domain+"_________________"+mr.KeywordSiteMatchInfo.Keyword])
			if err != nil {
				panic(err)
			}
		}
		tasks = [][]string{}

	}
}

func writeLine(taskR seoinfoQuery.MultiResult, wr csv.Writer, rank int) error {
	if taskR.SiteSeoInfo == nil || taskR.KeywordSiteMatchInfo == nil || taskR.KeywordMatchInfo == nil {
		return nil
	}
	err := wr.Write([]string{
		taskR.SiteSeoInfo.Domain,
		taskR.KeywordSiteMatchInfo.Keyword,
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
		strconv.Itoa(rank),
	})
	wr.Flush()
	return err
}
func combineTwoMap(dRMap map[string]int, dRMap2 map[string]int) map[string]int {
	if dRMap2 != nil {
		for d, r := range dRMap2 {
			if _, ok := dRMap[d]; !ok {
				dRMap[d] = r
			}
		}
	}

	return dRMap
}
