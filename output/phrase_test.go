package output

import (
	"fmt"
	"testing"

	"github.com/auho/go-nlp/output/component/dep"
	"github.com/auho/go-nlp/output/component/pos"
	"github.com/auho/go-nlp/output/component/tokenize"
)

func TestPhrase(t *testing.T) {
	/*
		东西不错，声音也不大，打扫也挺干净的，不错不错
		["东西", "不错", "，", "声音", "也", "不", "大", "，", "打扫", "也", "挺", "干净", "的", "，", "不错", "不错"]
		["NN", "VA", "PU", "NN", "AD", "AD", "VA", "PU", "VV", "AD", "AD", "VA", "SP", "PU", "VA", "VA"]
		[[2, "nsubj"], [0, "root"], [2, "punct"], [7, "nsubj"], [7, "advmod"], [7, "neg"], [2, "dep"], [2, "punct"], [12, "nsubj"], [12, "advmod"], [12, "advmod"], [2, "dep"], [12, "dep"], [2, "punct"], [2, "dep"], [15, "comod"]]
	*/

	_s := Sentence{
		s:   "东西不错，声音也不大，打扫也挺干净的，不错不错",
		tok: []tokenize.Tokenizes{{"东西", "不错", "，", "声音", "也", "不", "大", "，", "打扫", "也", "挺", "干净", "的", "，", "不错", "不错"}},
		pos: []pos.Poses{{"NN", "VA", "PU", "NN", "AD", "AD", "VA", "PU", "VV", "AD", "AD", "VA", "SP", "PU", "VA", "VA"}},
		dep: []dep.Deps{
			{
				dep.Dep{RootIndex: 1, Tag: "nsubj"},
				dep.Dep{RootIndex: 0, Tag: "root"},
				dep.Dep{RootIndex: 1, Tag: "punct"},
				dep.Dep{RootIndex: 6, Tag: "nsubj"},
				dep.Dep{RootIndex: 6, Tag: "advmod"},
				dep.Dep{RootIndex: 6, Tag: "neg"},
				dep.Dep{RootIndex: 1, Tag: "dep"},
				dep.Dep{RootIndex: 1, Tag: "punct"},
				dep.Dep{RootIndex: 11, Tag: "nsubj"},
				dep.Dep{RootIndex: 11, Tag: "advmod"},
				dep.Dep{RootIndex: 11, Tag: "advmod"},
				dep.Dep{RootIndex: 1, Tag: "dep"},
				dep.Dep{RootIndex: 11, Tag: "dep"},
				dep.Dep{RootIndex: 1, Tag: "punct"},
				dep.Dep{RootIndex: 1, Tag: "dep"},
				dep.Dep{RootIndex: 14, Tag: "comod"},
			},
		},
	}

	_ps := _s.ToPhrase()

	for _, s := range []string{`AD`, `AD AD`, `VA`, `AD( AD)* VA`} {
		_extractAd := NewExtractPos(s)
		for _, _p := range _ps {
			_r := _extractAd.Extract(_p)
			fmt.Println(_r)
		}
	}
}
