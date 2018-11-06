package syntax

import (
	"regexp"
	"strings"
)

type phrase struct {
	reg []*regexp.Regexp
}

type wordGroup struct {
	phrase     string // 匹配的短语
	vocabulary string // 匹配的短语 voc 组合
}

func NewPhrase() *phrase {
	return &phrase{}
}

func (p *phrase) MultipleMatchmakers(format []string) {
	for _, matchmaker := range format {
		p.Matchmaker(matchmaker)
	}
}

/*
 reg
	n a
	n v a

	"/n /a "
	"/n /v /a "
 */
func (p *phrase) Matchmaker(reg string) {
	reg = "/" + strings.Replace(reg, " ", " /", -1) + " "
	p.reg = append(p.reg, regexp.MustCompile(reg))
}

func (p *phrase) Seek(beau []string) []wordGroup {
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
	var wordGroups []wordGroup
	for _, reg := range p.reg {
		indexes := reg.FindAllStringIndex(vocsString, -1)

		for _, index := range indexes {
			wordGroups = append(wordGroups, p.propose(words, vocsString, index))
		}
	}

	return wordGroups
}

func (p *phrase) propose(words []string, vocsString string, indexes []int) wordGroup {
	c := wordGroup{}

	matchVoc := vocsString[indexes[0]:indexes[1]]
	startIndex := strings.Count(vocsString[0:indexes[0]], " ")
	offsetIndex := strings.Count(matchVoc, " ") + 1

	c.phrase = p.pickUpSlice(words, startIndex, startIndex+offsetIndex)
	c.vocabulary = strings.Replace(vocsString[indexes[0]:indexes[1]], "/", "", -1)

	return c
}

func (p *phrase) pickUpSlice(items []string, startIndex int, endIndex int) string {
	return strings.Join(items[startIndex:endIndex], "")
}
