package ctb

import (
	"github.com/auho/go-nlp/output/component/pos/tag"
)

const TagAD tag.Tag = "AD"   // adverb	副词
const TagAS tag.Tag = "AS"   // aspect marker	动态助词
const TagBA tag.Tag = "BA"   // bǎ in ba-construction	把字句
const TagCC tag.Tag = "CC"   // coordinating conjunction	并列连接词
const TagCD tag.Tag = "CD"   // cardinal number	概数词
const TagCS tag.Tag = "CS"   // subordinating conjunction	从属连词
const TagDEC tag.Tag = "DEC" // de as a complementizer or a nominalizer	补语成分“的”
const TagDEG tag.Tag = "DEG" // de as a genitive marker and an associative marker	属格“的”
const TagDER tag.Tag = "DER" // resultative de, de in V-de const and V-de-R	表结果的“得”
const TagDEV tag.Tag = "DEV" // manner de, de before VP	表方式的“地”
const TagDT tag.Tag = "DT"   // determiner	限定词
const TagETC tag.Tag = "ETC" // for words like “etc.”	表示省略
const TagEM tag.Tag = "EM"   // emoji	表情符
const TagFW tag.Tag = "FW"   // foreign words	外来语
const TagIC tag.Tag = "IC"   // incomplete component	不完整成分
const TagIJ tag.Tag = "IJ"   // interjection	句首感叹词
const TagJJ tag.Tag = "JJ"   // other noun-modifier	其他名词修饰语
const TagLB tag.Tag = "LB"   // bèi in long bei-const	长句式表被动
const TagLC tag.Tag = "LC"   // localizer	方位词
const TagM tag.Tag = "M"     // measure word	量词
const TagMSP tag.Tag = "MSP" // other particle	其他小品词
const TagNN tag.Tag = "NN"   // common noun	其他名词
const TagNOI tag.Tag = "NOI" // noise that characters are written in the wrong order	噪声
const TagNR tag.Tag = "NR"   // proper noun	专有名词
const TagNT tag.Tag = "NT"   // temporal noun	时间名词
const TagOD tag.Tag = "OD"   // ordinal number	序数词
const TagON tag.Tag = "ON"   // onomatopoeia	象声词
const TagP tag.Tag = "P"     // preposition e.g., “from” and “to”	介词
const TagPN tag.Tag = "PN"   // pronoun	代词
const TagPU tag.Tag = "PU"   // punctuation	标点符号
const TagSB tag.Tag = "SB"   // bèi in short bei-const	短句式表被动
const TagSP tag.Tag = "SP"   // sentence final particle	句末助词
const TagURL tag.Tag = "URL" // web address	网址
const TagVA tag.Tag = "VA"   // predicative adjective	表语形容词
const TagVC tag.Tag = "VC"   // copula, be words	系动词
const TagVE tag.Tag = "VE"   // yǒu as the main verb	动词有无
const TagVV tag.Tag = "VV"   // other verb	其他动词
