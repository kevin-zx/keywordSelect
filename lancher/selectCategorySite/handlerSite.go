package main

import (
	"bufio"
	"fmt"
	"github.com/kevin-zx/seotools/comm/site_base"
	"io"
	"jinzhunassist/domain"
	"os"
	"strings"
)

var f2s = []string{"seo网站优化", "seo分析工具", "seo营销", "seo搜索优化", "seo网络推广", "网站搜索排名优化", "搜索引擎seo", "东莞网站优化", "优化公司", "百度优化服务", "快速收录", "在线seo", "刷百度关键词排名", "提高百度排名", "自动seo", "优化方法", "威海seo", "seo攻略", "seo怎么做", "关键词优化李守洪排名大师", "seo案例分析", "排名seo", "seo培训教程", "关于seo", "培训seo", "昆明网站优化", "快速排名技术", "莆田seo", "seosem区别", "seo点击软件", "seo秘籍", "seo优化效果", "长沙seo顾问", "李守洪排名大师", "seo外链网站", "如何做好网站优化", "seo软件哪个好", "seo优化行业", "郑州seo培训", "seo营销是什么", "潮州seo", "seo查询工具", "引擎优化", "搜索引擎优化服务", "嘉兴网站优化", "泉州seo", "seo优化报价", "seo优化常识", "网站推广", "刷网站排名软件", "seo点击", "网络优化课程", "软件优化网站", "优化网站排名", "seo网站优化软件", "优化技术", "百度网站排名", "网络优化", "如何进行seo", "优化关键字", "如何提高百度权重", "网站优化软件", "seo书籍", "seo前景如何", "网站排名seo", "seo优化设置", "最好的seo学习网站", "优化是什么意思", "优化推广", "seo赚钱", "seo交流", "郴州网站优化", "帮站seo", "北京网站优化", "福州网站优化", "网站内容优化", "单页优化", "湖南网站seo", "seo优化方法", "seo关键词刷排名", "seo大神", "seo自然排名", "温州网站优化", "怎么提升网站的排名", "上海关键词优化", "杭州seo顾问", "网页优化", "盐城seo", "快速排名", "好搜优化", "seo免费诊断", "搜索引擎排行", "什么是优化", "seo专业培训", "网站免费优化", "草根seo", "怎么提升网站排名", "如何做seo推广", "网站建设优化", "黑帽seo教程", "刷网站排名", "怎么优化关键词排名", "网站seo优化公司", "包头seo", "如何做搜索引擎优化", "seo前景", "seo是什么意思", "网站权重", "seo统计", "seo网络公司", "怎么做seo", "网站推广与优化", "网站优化查询", "网站建设seo", "苏州seo培训", "seo刷词", "网站关键字优化", "seo菜鸟论坛", "seo优化教材", "seo关键词优化软件", "seo和sem的区别", "百度排名", "天津seo博客", "seo北京", "如何优化一个网站", "网站关键词优化排名", "刷排名软件", "刷排名网址", "怎么优化关键词", "如何网站推广", "百度关键词搜索", "上海seo服务", "怎么优化", "seo推广公司", "排名优化课程", "seo诊断分析工具", "搜索引擎排名", "滁州seo", "seo图片", "seo基础", "Seo刷排名", "关键词排名优化", "网络优化是什么", "关键词排名点击器", "百度快照优化", "怎么优化网站", "网站怎么优化", "长沙网站优化", "排名点击器", "上海seo优化", "seo优化知识", "如何做好网络推广", "站外seo", "网站的优化", "seo搜索引擎优化", "网络优化seo", "如何提高seo", "关键字密度", "如何做seo优化", "如何提高网站权重", "百度seo排名公司", "网站seo教程", "百度seo优化排名", "seo技巧", "seo是什么", "东莞seo培训", "seo名录", "百度搜索优化", "网奇seo", "seo引擎优化软件", "优化", "百度排名优化软件", "怎么做网站优化", "seo数据分析", "深圳seo公司", "安阳seo", "百度关键词优化", "seo优化策略", "如何做好seo", "单页面优化", "网络推广公司排名", "网站权重提升", "广州seo", "seo全攻略", "关键词的优化", "莱芜网站优化", "刷百度排名", "广州seo排名", "百度优化", "惊雷算法", "seo教学", "seo优化网络", "关键词密度", "网站关键词", "seo关键词软件", "seo研究中心怎么样", "seo教程", "提高网站权重", "搜索引擎网站推广", "杭州seo优化", "博客优化", "seo培训资料", "三明seo", "做网站seo", "免费seo", "搜索引擎优化工具", "seo关键词优化工具", "淮南seo赛雷猴", "优化一个网站", "seo首页优化", "关键词排名提升", "seo哪家好", "百度关键字", "seo关键词优化", "邵阳seo", "湖南网站优化", "网站优化公司", "站外优化", "大庆seo", "seo免费教程", "如何优化关键词", "网站推广公司", "邵阳网站优化", "关键词排名优化软件", "福州seo推广", "公司网站推广技巧", "seo", "湖南seo", "seo观察", "上海seo", "seo关键词优化排名", "seo分析", "北京seo培训", "seo接单", "怎么推广网站", "seo优化技巧", "优化点击软件", "搜索引擎关键词优化", "seo编辑培训", "南昌网站优化", "网站seo如何优化", "手机关键词排名", "什么是seo技术", "seo文案", "百度推广优化", "百度快照", "seo学习", "如何优化搜索引擎", "商丘网站优化", "随州seo", "网站seo技巧", "宁波seo", "关键词搜索排名", "seo站外优化", "seo方法", "网站优化方法", "实战seo培训", "如何优化网站", "电商seo", "福建seoseo8", "南阳seo", "seo在线培训机构", "seo顾问", "seo平台", "seo工作内容", "专业网站优化", "seo网络优化师", "seo搜索引擎", "网站优化排名", "关键词快速排名软件", "百度seo培训", "seo优化网", "快速排名首页", "seo排名优化软件", "中山seo", "太原seo", "seo服务", "站内优化", "seo优化怎么做", "网店优化推广", "长尾词快速排名", "SEO", "武汉seo服务", "怎么做seo推广", "seo网站优化方案", "关键词优化服务", "谷歌seo优化", "推广优化", "排名软件", "seo学院", "杭州seo公司", "扬州seo", "上海网站优化", "企业seo优化", "seo刷排名工具", "网络优化培训机构", "如何seo", "seo自动", "网络营销公司排名", "seo与sem", "百度seo优化", "seo推广优化", "排名优化软件", "快速提升排名", "无锡seo", "SEO测试", "百度seo排名点击软件", "网站seo方案", "外链seo", "站长seo工具", "快速seo优化", "百度seo关键词排名", "seo排名", "福州seo服务", "内链优化", "长沙网站seo", "seo发外链", "百合seo培训", "好搜seo软件", "seo黑帽技术", "seo排名点击器", "seo原理", "肇庆seo", "关键词排名如何提升", "alexa优化", "白帽seo技术", "长沙网站优化培训", "长沙seo", "灰帽seo", "广西seo", "搜索排名", "seo新人培训", "seo关键词快速排名", "嘉兴seo优化", "seo优化交流", "seo网站关键词优化", "seo推广教程", "seo关键词排名优化", "seo站群", "seo入门", "站长工具seo", "seo整站优化方案", "优化教程", "优化服务", "seo怎么样", "什么是搜索引擎优化", "江门seo", "最好的seo", "北京seo排名", "seo排名点击软件", "站群seo", "seo9", "百度seo教程", "企业网站优化", "昆明seo优化", "湘潭网站优化", "seo运营", "关键字搜索排名", "seo什么意思", "seo查询", "广东seo", "网站搜索排名", "手机端关键词快速排名", "seo快速排名软件", "seo案例", "关键字推广", "seo优化教程", "优化网站方法", "百度关键词优化工具", "移动seo", "高级seo", "合肥seo", "免费网站seo诊断", "搜索引擎怎么优化", "黑帽刷排名", "优化策略", "seo行业", "网络排名优化软件", "电子商务seo", "seo收费", "seo技术教程", "网络优化服务", "seo网页优化", "搜索关键词排名", "seosem", "网站关键词优化教程", "seo关键词工具", "seo网站", "seo教", "西安seo培训", "上海seo培训", "什么seo", "seo外包", "seo营销培训", "公司优化", "黑帽seo技术", "英文网站seo", "汕头网站优化", "seo如何优化", "网站优化", "seo优化培训", "网络推广", "网站关键词排名查询", "搜索引擎优化方法", "搜狗优化", "提升网站关键词排名", "商丘seo", "seo网络优化推广", "关键词策略", "seo关键词排名", "网站seo服务", "网站快速优化排名", "网址优化", "优化seo", "seo优化价格", "高质量外链", "关于seo优化", "百度推广seo", "seo咨询", "优化的意思", "南京seo顾问", "seo优化找狼雨", "刷百度排名软件", "上海seo学习", "什么是网站优化", "seo经验分享", "seo软件排行榜", "seo快速排名", "seo基础入门教程", "百度seo优化软件", "合肥关键词排名优化", "seo教程下载", "保定seo", "如何学习seo", "关键字排名", "seo好学吗", "百度网站优化软件", "南通seo", "百度seo排名优化软件", "seo排名工具", "厦门seo", "sem和seo区别", "企业网站推广技巧", "seo公司", "seo优化多少钱", "seo怎么赚钱", "seo优化网站", "什么是seo优化", "免费seo培训", "seo描述", "seo头条", "网站建设与优化", "seo优化是什么意思", "石家庄seo", "网站seo查询", "福州seo", "搜索引擎优化技术", "seo优化方案", "优化网站", "上海百度优化", "百度seo关键词优化", "博客seo优化", "台州seo", "电子商务网站seo", "seo资源", "seo优化点击软件", "seo优化视频", "英文网站优化", "百度seo软件", "重庆seo", "seo系统", "seo监测", "网站优化哪里好", "广西seo优化", "网络推广seo", "网页优化seo", "广州seo优化", "seo优化作用", "龙岩seo", "百度seo排名软件", "顶级seo", "搜爱seo", "优化外包", "刷百度关键词", "优化软件", "seo外链", "seo检测", "什么是seo", "手机网站seo", "网站关键词优化软件", "百度网站优化", "seo诊断", "网页搜索优化", "seo专员", "seo软件工具", "企业seo", "泰安seo", "网站排名", "seo优化师", "刷排名网站", "绍兴seo", "关键词排名点击", "seo引擎优化", "快排优化", "杭州seo培训", "seo优化论坛", "博客seo", "seo培训班", "网页seo", "网站排名优化方法", "seo优化技术", "google关键词优化", "推广seo", "seo导航", "网站优化课程", "baidu优化", "怎么提高关键词排名", "seo推广", "seo资讯", "深圳seo教程", "seo整站优化", "黑帽seo", "seo网站推广", "公司网站优化", "seo教程网", "seo的优化", "无锡seo优化", "百度seo优化公司", "山西seo优化", "天津seo优化", "搜索引擎排名优化", "德阳seo", "网络seo", "网站排名优化软件", "如何做网站推广优化", "海南seo", "新站排名", "网络推广优化", "搜狗seo", "seo网站内部优化", "长春网站seo", "青岛seo培训", "什么是优化推广", "seo方案", "百度关键词排名", "网站优化方式", "北京seo服务", "百度网页快照", "优化培训班", "百度优化公司", "关键词优化软件", "seo培训课程", "如何提升网站的排名", "长春网站优化", "seo代理", "新站seo", "刑天seo", "网站如何优化", "网站关键词推广", "郑州seo顾问", "网站搜索引擎优化", "无锡seo公司", "上海seo顾问", "seo快速优化", "百度seo关键词", "网络优化公司", "疯狂seo", "seo研究中心", "网站排名优化培训", "seo新手入门教程", "简单seo", "百度seo排名优化", "seo排名培训", "seo站长工具", "移动端seo", "seo排名优化", "台州网站优化", "seo课程", "秦皇岛网站优化", "南京seo优化", "semseo区别", "优化快速排名", "网站如何优化排名", "关键词优化", "怎么seo优化", "免费seo诊断", "seo研究", "网站优化方案", "seo教学视频", "网页优化技巧", "杭州seo", "新站优化", "常州seo", "seo管理", "seo推广软件", "seo排名点击", "seo视频", "网站首页被k", "seo技术交流", "seo怎么优化", "宝鸡seo", "seo实战课程", "seo总结", "快排seo软件", "太原seo优化", "关键词优化报价", "seo基础知识", "优化怎么做", "医疗seo", "网站推广的方式", "seo基础入门", "黑帽seo软件", "seo论坛", "武汉seo培训", "seo营销工具", "seo策略", "刷关键词排名", "蚌埠seo", "seo优化", "seo课程培训", "seo免费培训", "seo排名优化培训", "seo刷排", "网站站内优化", "seo门户网", "优化排名软件", "湖南seo优化", "seo和sem", "手机排名seo", "武汉seo顾问", "seo视频培训", "百度优化排名", "百度排名点击软件", "整站优化", "旺道seo", "莱芜seo", "宁德seo", "网站关键词如何优化", "网络优化推广", "怎样做seo", "太原网站优化", "关键词排名优化工具", "云南seo", "企业网站排名优化", "网站排名优化", "seo实战培训", "深圳seo优化", "网站推广技巧", "北京seo顾问", "百度优化关键词", "阳江seo", "快速排名优化", "seo流量", "SEO优化", "seo监控", "天津seo诊断", "关键词自然排名优化", "seo免费培训教程", "九江seo", "北京seo优化", "北京seo公司", "seo博客优化", "seo团队", "网站seo怎么做", "如何推广网站", "如何做seo", "seo与sem的区别", "seo常用工具", "长沙百度优化", "淘宝seo优化", "seo自学网", "惠州seo", "seo关键词", "网站优化排名软件", "网站不收录", "图片seo", "百度优化排名软件", "seo怎么学", "seo优化公司", "seo排名查询", "搜索词排名", "如何提高百度排名", "关键字怎么优化", "seo怎么优化关键词", "seo培训多少钱", "seo辉煌电商平台", "seo公司排名", "网站优化培训", "网络营销优化", "关键字排名优化", "seo关键词优化外包", "seo优化建议", "搜索引擎排行榜", "宁波seo优化", "快速排名软件", "seo内部优化", "唐山seo", "刷网站权重", "网站页面优化", "企业站seo", "好搜seo", "深圳seo", "荆州seo", "网页搜索引擎优化", "企业网络优化方案", "优化培训", "搜索引擎优化", "南宁seo优化", "seo技术", "百度seo排名", "网站优化seo", "搜索排名优化", "北京seo软件", "网站关键词排名优化", "网站排名工具", "免费seo网站诊断", "seo云优化", "上海网站排名优化", "seo网", "山西seo", "搜索网站排名", "网站搜索引擎排名", "张岩seo", "seo学习论坛", "关键词排名方案", "关键词优化排名软件", "快速排名seo8", "seo网站排名优化", "贵阳seo", "301转向", "seo外链怎么发", "如何快速收录", "seo文章优化", "搜索引擎优化排名", "百度关键词推广", "网站死链查询", "站长工具seo推广", "seocnm", "seo报价", "优化排名工具", "快照优化", "seo网络培训", "网站排名提高", "黑客seo", "搜索引擎优化seo", "龙岩网站优化", "网站优化教程", "seo基础培训", "深圳seo培训", "怎么做好seo", "站长seo", "网站建设", "网站推广优化", "百度seo", "seo工作", "seo实战", "山东seo", "优化网", "官网优化", "seo搜索", "医院seo", "金华seo", "关键词优化排名", "外链优化", "seo搜索推广", "软文seo", "怎么做好seo优化", "seo思维", "SEO优化课程", "seo搜索优化软件", "说说seo", "seo优化的网站", "网站排名软件", "哈尔滨seo", "seo培训网", "seo要学多久", "提高seo", "排名工具", "百度关键词", "怎么seo", "上海seo公司", "优化营销", "网站怎么优化推广", "嘉兴seo", "长沙seo优化", "百度seo建议", "杭州网站优化", "官网seo", "网络优化是做什么的", "seo点击工具", "企业seo培训", "外链建设", "实战seo", "看seo", "提升网站排名", "seo标签", "seo站长", "网络营销", "页面seo优化", "淮南seo", "网站代码优化", "武汉seo论坛", "提高网站排名", "seo网络优化", "seo培训机构", "百度seo网站优化", "苏州旺道seo", "seo外包公司", "浙江seo", "长春seo", "优化诊断", "网址排名", "网站seo优化", "seo优化顾问", "页面优化", "论坛优化seo", "惠州seo优化", "seo简单", "seo知识", "seo推广方案", "优化百度", "长沙seo培训", "站长排名", "百度收录", "温州seo", "seo建站", "网站seo培训", "seo点击器", "seo公司哪家好", "网站如何做优化", "百度不收录", "seo在线培训", "网站seo顾问", "seo电子书", "网站seo", "seo站内优化", "常德seo", "seo标题优化", "狼雨seo", "提升关键词排名", "泰州seo", "高质量外链资源", "关键词seo优化", "排名提升", "死链接", "网站优化平台", "江门网站优化", "网站优化推广", "快速网站排名", "搜索引擎优化师", "关键词优化公司", "关键词优化技巧", "怎么优化自己网站", "网站权重优化", "汕头seo优化", "seo自动推广工具", "如何自学seo", "seo权重", "seo优化软件", "seo工具哪个好", "零基础学seo", "seo赚钱培训", "优化工具", "西风seo", "湘潭seo", "东莞seo", "网站推广排名", "什么叫网站优化", "seo专家", "百度搜索排名", "关键字优化", "搜索优化", "搜狗搜索引擎优化", "seo代码优化", "seo菜鸟", "百度seo公司", "seo项目", "整站排名优化", "锚文本", "网站搜索优化", "seo提交", "快速优化", "seo推广工具", "镇江seo", "医院网站优化", "seo优化服务", "如何提高关键词排名", "seo技术培训", "网站优化工具", "seo排行", "百度快速排名软件", "关键词优化方法", "企业网站seo", "seo快速优化软件", "如何seo推广", "seo新闻", "济南seo培训", "长春seo优化", "优化关键词排名", "温州seo优化", "优化工作", "网站建设排名", "兰州seo", "seo学习网", "北京seo优化公司", "seo入门教程", "移动网站优化", "搜索引擎优化教程", "seo外链工具", "seo网站建设", "学习seo", "seo在线优化工具", "白帽seo", "站内seo", "关键词优化分析", "关键词排名", "关键词seo排名", "网站死链", "站点排名", "百度网站收录", "电商网站seo", "搜索引擎优化学习", "深圳网站优化", "如何优化关键词排名", "重庆seo优化", "seo服务公司", "seo门户", "seo顾问服务", "seo在线优化", "seo优化关键词", "网站排名提升", "黑帽seo培训", "seo标题", "网站优化建议", "seo自学", "seo外链建设", "江苏seo", "seo网站诊断", "北京seo外包", "百度关键词优化软件", "seo博客", "网站优化关键词", "seo工具", "推广网站排名", "汕头seo", "学seo", "seo优化是什么", "网站关键词排名", "新网站seo", "北京seo", "搜索引擎优化培训", "百度seo排名点击器", "专业seo优化", "seo快排", "seo培训", "优化网络", "seo软件", "黑帽技术", "seo关键词推广", "百度关键字优化", "企业网站怎么优化", "石家庄seo培训", "九成seo", "内容优化", "北京seo方法", "网站seo诊断", "怎样优化网站", "seo培训公司", "优化排名", "百度关键词seo", "排名优化", "网站快速排名软件", "排名优化公司", "seo指令", "网站关键词没有排名", "网站首页优化", "网站诊断", "如何优化网站排名", "seo排名软件", "企业网站如何优化", "seo外链推广", "广安seo", "搜索引擎优化技巧", "洛阳seo", "百度排名优化", "seo基础教程", "百度搜索引擎优化", "权重优化", "锚文本链接", "建站seo", "seo优化页面", "佛山seo", "百度优化软件", "seo系统培训", "驻马店seo", "seo优化工具", "seo术语", "seo教程视频", "seo视频教程", "网站关键词优化"}

func main() {
	d1, err := os.Open("data/d2.csv")
	if err != nil {
		panic(err)
	}
	defer d1.Close()

	br := bufio.NewReader(d1)
	for {
		line, err := br.ReadString('\n')
		m := false
		for _, f2 := range f2s {
			if strings.Contains(line, f2) {
				m = true
				break
			}
		}
		if !m {
			continue
		}
		ps := strings.Split(line, ",")
		siteDomain := ps[0]
		handler(siteDomain)
		if err == io.EOF {
			break
		}
	}

}
func handler(siteDomain string) {
	si, err := site_base.ParseWebSeoFromUrl("http://" + siteDomain)
	if err != nil {
		return
	}
	m := false
	for _, f2 := range f2s {
		if strings.Contains(si.Keywords, f2) {
			m = true
			break
		}
	}
	if m {
		si5118, err := domain.GetDomainInfo(siteDomain)
		if err != nil {
			return
		}

		fmt.Println(si.RealUrl, formatString2(si.Title), si5118.KeywordCount, si5118.PcPvSum, si5118.MobilePvSum)
	}

}

func formatString2(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "\"", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "'", "", -1)
	return s
}
