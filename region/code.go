package region

//go:generate enumer --type=Code --linecomment --json --extramethod --output=code_enum.go
type Code int

const (
	// 厄瓜多尔
	CodeEC Code = iota + 1 // EC
	// 巴西
	CodeBR // BR
	// 墨西哥
	CodeMX // MX
	// 哥伦比亚
	CodeCO // CO
	// 印度尼西亚
	CodeID // ID
	// 菲律宾
	CodePH // PH
	// 秘鲁
	CodePE // PE
	// 埃及
	CodeEG // EG
	// 土耳其
	CodeTR // TR
	// 阿根廷
	CodeAR // AR
	// 智利
	CodeCL // CL
	// 委内瑞拉
	CodeVE // VE
	// 玻利维亚
	CodeBO // BO
	// 印度
	CodeIN // IN
	// 日本
	CodeJP // JP
	// 泰国
	CodeTH // TH
	// 马来西亚
	CodeMY // MY
	// 越南
	CodeVN // VN
	// 俄罗斯
	CodeRU // RU
	// 新加坡
	CodeSG // SG
	// 孟加拉国
	CodeBD // BD
	// 韩国
	CodeKR // KR
	// 阿塞拜疆
	CodeAZ // AZ
	// 巴拿马
	CodePA // PA
	// 约旦
	CodeJO // JO
	// 阿曼苏丹国
	CodeOM // OM
	// 哈萨克斯坦
	CodeKZ // KZ
	// 沙特阿拉伯
	CodeSA // SA
	// 英国
	CodeUK // UK
	// 多米尼加共和国
	CodeDO // DO
	// 福克兰群岛
	CodeFK // FK
	// 法属圭亚那
	CodeGF // GF
	// 圭亚那
	CodeGY // GY
	// 巴拉圭
	CodePY // PY
	// 苏里南
	CodeSR // SR
	// 乌拉圭
	CodeUY // UY
	// 加拿大
	CodeCA // CA
	// 美国
	CodeUS // US
	// 伯利兹
	CodeBZ // BZ
	// 哥斯达黎加
	CodeCR // CR
	// 古巴
	CodeCU // CU
	// 萨尔瓦多
	CodeSV // SV
	// 危地马拉
	CodeGT // GT
	// 海地
	CodeHT // HT
	// 洪都拉斯
	CodeHN // HN
	// 牙买加
	CodeJM // JM
	// 尼加拉瓜
	CodeNI // NI
	// 波多黎各
	CodePR // PR
	// 特立尼达和多巴哥
	CodeTT // TT
	// 阿尔巴尼亚
	CodeAL // AL
	// 安道尔
	CodeAD // AD
	// 奥地利
	CodeAT // AT
	// 白俄罗斯
	CodeBY // BY
	// 比利时
	CodeBE // BE
	// 波斯尼亚和黑塞哥维那
	CodeBA // BA
	// 保加利亚
	CodeBG // BG
	// 克罗地亚
	CodeHR // HR
	// 塞浦路斯
	CodeCY // CY
	// 捷克
	CodeCZ // CZ
	// 丹麦
	CodeDK // DK
	// 爱沙尼亚
	CodeEE // EE
	// 法罗群岛
	CodeFO // FO
	// 芬兰
	CodeFI // FI
	// 法国
	CodeFR // FR
	// 德国
	CodeDE // DE
	// 希腊
	CodeGR // GR
	// 中国台湾
	CodeTW // TW
	// 阿富汗
	CodeAF // AF
	// 伊朗
	CodeIR // IR
	// 伊拉克
	CodeIQ // IQ
	// 科威特
	CodeKW // KW
	// 吉尔吉斯斯坦
	CodeKG // KG
	// 黎巴嫩
	CodeLB // LB
	// 马尔代夫
	CodeMV // MV
	// 蒙古
	CodeMN // MN
	// 尼泊尔
	CodeNP // NP
	// 乌兹别克斯坦
	CodeUZ // UZ
	// 也门
	CodeYE // YE
	// 南非
	CodeZA // ZA
	// 以色列
	CodeIL // IL
	// 澳大利亚
	CodeAU // AU
	// 新西兰
	CodeNZ // NZ
	// 意大利
	CodeIT // IT
	// 尼日利亚
	CodeNG // NG
	// 摩洛哥
	CodeMA // MA
	// 斯里兰卡
	CodeLK // LK
	// 美国虚拟卡
	CodeUSV // USV
	// 几内亚
	CodeGN // GN
	// 中国香港
	CodeHK // HK
	// 阿联酋
	CodeAE // AE
	// 波兰
	CodePL // PL
	// 摩纳哥
	CodeMC // MC
	// 爱尔兰
	CodeIE // IE
	// 老挝
	CodeLA // LA
	// 叙利亚
	CodeSY // SY
	// 拉脱维亚
	CodeLV // LV
	// 荷兰
	CodeNL // NL
	// 葡萄牙
	CodePT // PT
	// 亚美尼亚
	CodeAM // AM
	// 西班牙
	CodeES // ES
	// 匈牙利
	CodeHU // HU
	// 斯洛伐克
	CodeSK // SK
	// 斯洛文尼亚
	CodeSI // SI
	// 立陶宛
	CodeLT // LT
	// 加蓬
	CodeGA // GA
	// 東帝汶
	CodeTL // TL
	// 塞尔维亚
	CodeRS // RS
	// 挪威
	CodeNO // NO
	// 冰島
	CodeIS // IS
	// 卢森堡
	CodeLU // LU
	// 黑山共和国
	CodeME // ME
	// 土库曼斯坦
	CodeTM // TM
	// 阿尔及利亚
	CodeDZ // DZ
	// 瑞典
	CodeSE // SE
	// 罗马尼亚
	CodeRO // RO
	// 中国澳门
	CodeMO // MO
	// 马里
	CodeML // ML
	// 巴林
	CodeBH // BH
	// 乌克兰
	CodeUA // UA
	// 美国（物理)
	CodeUSP // USP
	// 安哥拉
	CodeAO // AO
	// 肯尼亚
	CodeKE // KE
	// 乍得
	CodeTD // TD
	// 蒙特塞拉特
	CodeMS // MS
	// 刚果
	CodeCG // CG
	// 突尼斯
	CodeTN // TN
	// 莫桑比克
	CodeMZ // MZ
	// 塞内加尔
	CodeSN // SN
	// 巴布亞新幾內亞
	CodePG // PG
	// 納米比亞
	CodeNA // NA
	// 埃塞俄比亚
	CodeET // ET
	// 利比亚
	CodeLY // LY
	// 尼日尔
	CodeNE // NE
	// 巴基斯坦
	CodePK // PK
	// 萊索托
	CodeLS // LS
	// 乌干达
	CodeUG // UG
	// 贊比亞
	CodeZM // ZM
	// 马达加斯加
	CodeMG // MG
	// 缅甸
	CodeMM // MM
	// 多哥
	CodeTG // TG
	// 加纳
	CodeGH // GH
	// 博茨瓦納
	CodeBW // BW
	// 冈比亚
	CodeGM // GM
	// 巴哈馬
	CodeBS // BS
	// 斐濟
	CodeFJ // FJ
	// 索馬里
	CodeSO // SO
	// 貝寧
	CodeBJ // BJ
	// 科特迪瓦
	CodeCI // CI
	// 柬埔寨
	CodeKH // KH
	// 喀麦隆
	CodeCM // CM
	// 布隆迪
	CodeBI // BI
	// 利比里亞
	CodeLR // LR
	// 几内亚比绍
	CodeGW // GW
	// 北馬其頓
	CodeMK // MK
	// 阿魯巴島
	CodeAW // AW
	// 科摩羅
	CodeKM // KM
	// 津巴布韋
	CodeZW // ZW
	// 盧旺達
	CodeRW // RW
	// 圣卢西亚
	CodeLC // LC
	// 馬拉維
	CodeMW // MW
	// 毛里求斯
	CodeMU // MU
	// 毛里塔尼亞
	CodeMR // MR
	// 南蘇丹
	CodeSS // SS
	// 赤道几内亚
	CodeGQ // GQ
	// 中非共和國
	CodeCF // CF
	// 塞拉利昂
	CodeSL // SL
	// 斯威士兰
	CodeSZ // SZ
	// 安圭拉島
	CodeAI // AI
	// 瑞士
	CodeCH // CH
	// 聖多美和普林西比
	CodeST // ST
	// 巴巴多斯
	CodeBB // BB
	// 塞舌爾共和國
	CodeSC // SC
	// 坦桑尼亚
	CodeTZ // TZ
	// 圣基茨和尼维斯
	CodeKN // KN
	// 新喀里多尼亚
	CodeNC // NC
	// 塔吉克斯坦
	CodeTJ // TJ
	// 摩尔多瓦
	CodeMD // MD
	// 格鲁吉亚
	CodeGE // GE
	// 布基纳法索
	CodeBF // BF
	// 中国
	CodeCN // CN
	// 卡塔尔
	CodeQA // QA
	// 文莱
	CodeBN // BN
	// 瓜德罗普岛
	CodeGP // GP
	// 厄立特里亚
	CodeER // ER
	// 安提瓜和巴布达
	CodeAG // AG
	// 吉布地
	CodeDJ // DJ
	// 开曼群岛
	CodeKY // KY
	// 格林納達
	CodeGD // GD
	// 佛得角
	CodeCV // CV
	// 圣文森特和格林纳丁斯
	CodeVC // VC
	// 多米尼加
	CodeDM // DM
	// 所罗门群岛
	CodeSB // SB
	// 汤加王国
	CodeTO // TO
	// 不丹
	CodeBT // BT
	// 民主刚果（刚果金）
	CodeCD // CD
	// 列支敦士登
	CodeLI // LI
	// 留尼汪
	CodeRE // RE
	// 萨摩亚
	CodeWS // WS
	// 圣马丁
	CodeSX // SX
)
