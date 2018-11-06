package syntax

import (
	"fmt"
	"regexp"
	"strings"
)

type headAndTailLimited struct {
	reg []*regexp.Regexp
}

type couple struct {
	phrase     string // 匹配的短语
	head       string // 匹配的头
	tail       string // 匹配的尾
	vocabulary string // 匹配的短语 voc 组合
	pattern    string // 不包含头尾的 pattern
	patternVoc string // 不包含头尾的 pattern voc
}

func NewHeadAndTailLimited() *headAndTailLimited {
	return &headAndTailLimited{}
}

func (hatl *headAndTailLimited) MultipleMatchmakers(format [][]string) {
	for _, matchmaker := range format {
		hatl.Matchmaker(matchmaker[0], matchmaker[1], matchmaker[2])
	}
}

/*
 headRegexp n|vn
 tailRegexp a|aj
 */
func (hatl *headAndTailLimited) Matchmaker(headRegexp string, tailRegexp string, step string) {
	regexpFormat := `(?P<head>/(?:%s)\s)(?:/\w+\s){0,%s}?(?P<tail>/(?:%s)\s)`

	hatl.reg = append(hatl.reg, regexp.MustCompile(fmt.Sprintf(regexpFormat, headRegexp, step, tailRegexp)))
}

//func (hatl *headAndTailLimited) BlindDate(beau [] string, format [][]string) []couple {
//	hatl.MultipleMatchmakers(format)
//
//	return hatl.Seek(beau)
//}
//
//func (hatl *headAndTailLimited) Pursue(beau []string, headRegexp string, tailRegexp string, step string) []couple {
//	hatl.Matchmaker(headRegexp, tailRegexp, step)
//	return hatl.Seek(beau)
//}

func (hatl *headAndTailLimited) Seek(beau []string) []couple {
	beauLen := len(beau)

	wv := make([]string, 0, 2)
	words := make([]string, 0, beauLen)
	vocs := make([]string, 0, beauLen)

	for _, item := range beau {
		wv = strings.Split(item, "/")
		words = append(words, wv[0])
		vocs = append(vocs, "/"+wv[1])
	}

	vocsString := strings.Join(vocs, " ") + " "
	var couples []couple
	for _, reg := range hatl.reg {
		subMatchIndexes := reg.FindAllStringSubmatchIndex(vocsString, -1)

		for _, subMatchIndex := range subMatchIndexes {
			couples = append(couples, hatl.propose(words, vocsString, subMatchIndex))
		}
	}

	return couples
}

/*
voc slice 和 words slice 位置索引一一对应，根据匹配索引结果，提取最后的信息
计算 start 之前的空格数量、offset 的空格数量，根据空格数量，把 voc string 的匹配索引，转换为 voc slice 的索引

匹配索引结果
sub match indexes
[[0 1 2 3 4 5]]
0 匹配文本开始位置
1 匹配文本结束位置
2 匹配头开始位置
3 匹配头结束位置
4 匹配尾开始位置
5 匹配尾结束位置

分组名称 group names
[0 1 2]
0 为空字符串表示匹配全部
1 第一个分组名称（头）
2 第二个分组名称（尾）
 */
func (hatl *headAndTailLimited) propose(words []string, vocsString string, indexes []int) couple {
	c := couple{}

	matchVoc := vocsString[indexes[0]:indexes[1]]
	startIndex := strings.Count(vocsString[0:indexes[0]], " ")
	offsetIndex := strings.Count(matchVoc, " ") + 1

	c.phrase = hatl.pickUpSlice(words, startIndex, startIndex+offsetIndex)
	c.head = words[startIndex]
	c.tail = words[startIndex+offsetIndex-1]

	c.vocabulary = strings.Replace(vocsString[indexes[0]:indexes[1]], "/", "", -1)
	c.pattern = hatl.pickUpSlice(words, startIndex+1, startIndex+offsetIndex-1)
	c.patternVoc = strings.Replace(vocsString[indexes[3]:indexes[4]], "/", "", -1)

	return c
}

func (hatl *headAndTailLimited) pickUpSlice(items []string, startIndex int, endIndex int) string {
	return strings.Join(items[startIndex:endIndex], "")
}
