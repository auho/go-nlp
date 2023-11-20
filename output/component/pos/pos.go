package pos

import (
	"strings"

	"github.com/auho/go-nlp/output/component/dep/tag"
)

// Part-of-Speech tagging、POS
// 词性标注

type Poses []tag.Tag

func (p Poses) ToSyntax() string {
	var s []string
	for _, _tag := range p {
		s = append(s, string(_tag))
	}

	return strings.Join(s, " ")
}
