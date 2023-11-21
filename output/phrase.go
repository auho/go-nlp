package output

import (
	"regexp"
	"strings"

	"github.com/auho/go-nlp/output/component/dep"
	depTag "github.com/auho/go-nlp/output/component/dep/tag"
	"github.com/auho/go-nlp/output/component/pos"
	"github.com/auho/go-nlp/output/component/tokenize"
)

type Phrases []Phrase

type Phrase struct {
	s   string
	tok tokenize.Tokenizes
	pos pos.Poses
	dep dep.Deps

	posSyntax string
}

func (p *Phrase) ToPhrases() Phrases {
	var ps Phrases

	section := p.dep.ToSection()
	for _, _ss := range section {
		_p := Phrase{
			tok: p.tok[_ss.OriginBeginIndex : _ss.OriginEndIndex+1],
			pos: p.pos[_ss.OriginBeginIndex : _ss.OriginEndIndex+1],
			dep: _ss.ToDeps(),
		}

		ps = append(ps, buildPhrase(_p))
	}

	return ps
}

func buildPhrase(p Phrase) Phrase {
	p.s = strings.Join(p.tok, "")
	p.posSyntax = p.pos.ToSyntax()

	return p
}

type ExtractPos struct {
	re *regexp.Regexp
}

func NewExtractPos(str string) *ExtractPos {
	return &ExtractPos{
		// 前面增加空格，方便后续匹配（简化 regexp expression ）
		//  A B CD EFG
		re: regexp.MustCompile(" " + str + `\b`),
	}
}

func (e *ExtractPos) Extract(p Phrase) []ExtractResult {
	var er []ExtractResult

	_ps := p.ToPhrases()
	for _, _p := range _ps {
		er = append(er, e.extractPhrase(_p)...)
	}

	return er
}

func (e *ExtractPos) extractPhrase(p Phrase) []ExtractResult {
	var er []ExtractResult

	_psLen := len(p.posSyntax)
	res := e.re.FindAllStringIndex(" "+p.posSyntax+" ", -1) // 前后增加空格
	for _, _res := range res {
		var _er ExtractResult

		_start := _res[0]    // 包含（regexp expression 包含前缀空格，匹配时 pos syntax 被增加空格，不用修正位置）
		_stop := _res[1] - 2 // 不包含（regexp expression 包含后缀空格，匹配时 pos syntax 被增加空格，去除被匹配的后缀的空格，修正位置）

		if _stop >= _psLen {
			_stop = _psLen
		}

		_indexStart := strings.Count(p.posSyntax[0:_start], " ")
		_indexEnd := _indexStart + strings.Count(p.posSyntax[_start:_stop], " ") + 1

		_er.Tok = p.tok[_indexStart:_indexEnd]
		_er.Phrase = strings.Join(_er.Tok, "")
		_er.Pos = p.pos[_indexStart:_indexEnd]
		_er.PosSyntax = _er.Pos.ToSyntax()

		er = append(er, _er)
	}

	return er
}

type ExtractResult struct {
	Phrase    string
	Tok       tokenize.Tokenizes
	Pos       pos.Poses
	PosSyntax string
}

type ExtractDepResult struct {
	ExtractResult
	DepTag depTag.Tag
}
