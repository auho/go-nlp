package pos

import (
	"regexp"
)

type Extract struct {
	re *regexp.Regexp
}

func NewExtract(str string) *Extract {
	return &Extract{
		re: regexp.MustCompile(str),
	}
}

func (e *Extract) Extract(s string) {

}
