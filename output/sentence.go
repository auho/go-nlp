package output

import (
	"github.com/auho/go-nlp/output/component/dep"
	"github.com/auho/go-nlp/output/component/pos"
	"github.com/auho/go-nlp/output/component/tokenize"
)

/*
东西不错，声音也不大，打扫也挺干净的，不错不错
["东西", "不错", "，", "声音", "也", "不", "大", "，", "打扫", "也", "挺", "干净", "的", "，", "不错", "不错"]
["NN", "VA", "PU", "NN", "AD", "AD", "VA", "PU", "VV", "AD", "AD", "VA", "SP", "PU", "VA", "VA"]
[[2, "nsubj"], [0, "root"], [2, "punct"], [7, "nsubj"], [7, "advmod"], [7, "neg"], [2, "dep"], [2, "punct"], [12, "nsubj"], [12, "advmod"], [12, "advmod"], [2, "dep"], [12, "dep"], [2, "punct"], [2, "dep"], [15, "comod"]]
*/
type Sentence struct {
	s   string
	tok []tokenize.Tokenizes
	pos []pos.Poses
	dep []dep.Deps
}

func (s *Sentence) ToPhrase() Phrases {
	var ps []Phrase
	for i := range s.tok {
		_p := Phrase{
			tok: s.tok[i],
			pos: s.pos[i],
			dep: s.dep[i],
		}
		ps = append(ps, buildPhrase(_p))
	}

	return ps
}
