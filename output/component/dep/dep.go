package dep

import (
	"github.com/auho/go-nlp/output/component/dep/tag"
	"github.com/auho/go-nlp/output/component/dep/tag/sd"
)

// Dependency Parsing、DEP
// 依存句法分析

type Sections struct {
	OriginBeginIndex int
	OriginEndIndex   int
	Sections         []Section
}

func (s *Sections) ToDeps() Deps {
	var _deps Deps
	for _, _s := range s.Sections {
		_deps = append(_deps, _s.ToDep())
	}

	return _deps
}

type Section struct {
	OriginIndex     int     // origin index
	OriginRootIndex int     // origin 中心词的下标，从 0 开始
	RootIndex       int     // new 中心词的下标，从 0 开始
	Tag             tag.Tag // 与中心词的依存关系
}

func (s *Section) ToDep() Dep {
	return Dep{
		RootIndex: s.RootIndex,
		Tag:       s.Tag,
	}
}

type Dep struct {
	RootIndex int     // 中心词的下标，从 0 开始
	Tag       tag.Tag // 与中心词的依存关系
}

func (d *Dep) ToSection() Section {
	return Section{
		OriginRootIndex: d.RootIndex,
		Tag:             d.Tag,
	}
}

type Deps []Dep

// ToSection
// [][2]int => {{start index , end index}}
func (d Deps) ToSection() []Sections {
	var sss []Sections

	_rootIndex := -1 // 从 0 开始
	for i, _dep := range d {
		if _dep.Tag == sd.TagRoot {
			_rootIndex = i
		}
	}

	_isStart := false
	_preRootIndex := -1
	_sectionIndex := -1
	_sectionRootIndex := -1

	var _ss Sections
	for _index, _dep := range d {
		if _dep.Tag == sd.TagPunct {
			continue
		}

		_sectionIndex += 1
		_indexDep := _dep

		// root node or new root node
		if _index == _rootIndex || (_dep.RootIndex == _rootIndex && _dep.Tag == sd.TagDep) {
			_indexDep.RootIndex = _index
			_indexDep.Tag = sd.TagRoot

		}

		// end section
		if _isStart && _indexDep.RootIndex != _preRootIndex {
			_isStart = false
			sss = append(sss, _ss)
			_sectionIndex = 0
			_ss = Sections{}
		}

		// start section
		if !_isStart {
			_isStart = true
			_preRootIndex = _indexDep.RootIndex
			_sectionRootIndex = _preRootIndex - _index
			_ss.OriginBeginIndex = _index
		}

		_s := _indexDep.ToSection()
		_s.OriginIndex = _index
		_s.RootIndex = _sectionRootIndex

		_ss.OriginEndIndex = _index
		_ss.Sections = append(_ss.Sections, _s)
	}

	if len(_ss.Sections) > 0 {
		sss = append(sss, _ss)
	}

	return sss
}
