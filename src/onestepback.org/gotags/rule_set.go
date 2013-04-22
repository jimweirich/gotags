package main

import (
	"strings"
	"regexp"
)

type Rule struct {
	Pattern  *regexp.Regexp
	TagIndex int
	DefIndex int
	IsMulti  bool
}

func NewRule(pattern string, tagIndex, defIndex int, isMulti bool) *Rule {
	result := Rule {}
	result.Init(pattern, tagIndex, defIndex, isMulti)
	return &result
}

func (self *Rule) Init(pattern string, tagIndex, defIndex int, isMulti bool) {
	self.Pattern = regexp.MustCompile(pattern)
	self.TagIndex = tagIndex
	self.DefIndex = defIndex
	self.IsMulti  = isMulti
}

func (self *Rule) Match(data string) (string, string) {
	matches := self.Pattern.FindStringSubmatch(data)
	if len(matches) > 0 {
		return matches[self.TagIndex], matches[self.DefIndex]
	}
	return "", ""
}

func (self *Rule) Apply(tag *Tag, data string, lineCount, byteCount int) bool {
	tagname, defstring := self.Match(data)
	if tagname != "" {
		if self.IsMulti {
			self.AddMulti(tag, tagname, defstring, lineCount, byteCount)
		} else {
			tag.Add(tagname, defstring, lineCount, byteCount)
		}
		return true
	}
	return false
}

func (self *Rule) AddMulti(tag *Tag, tagname, defstring string, lineCount, byteCount int) {
	for _, name := range strings.Split(tagname, ",") {
		name = strings.Trim(name, " *:")
		tag.Add(name, defstring, lineCount, byteCount)
	}
}

func rules() []*Rule {
	rs := make([]*Rule, 5)
	rs[0] = NewRule("^[ \t]*(class|module)[ \t]+([^:]+::)*([A-Z][A-Za-z0-9_]*)", 3, 0, false)
	rs[1] = NewRule("^[ \t]*def[ \t]+((self\\.)?[a-z0-9_]+(!?)?)", 1, 0, false)
	rs[2] = NewRule("^[ \t]*([A-Z][A-Za-z0-9_]*)[ \t]*=", 1, 0, false)
	rs[3] = NewRule("^[ \t]*attr_(reader|writer|accessor)[ \t]+([:a-z0-9_, ]+)", 2, 0, true)
	rs[4] = NewRule("^[ \t]*alias[ \t]+:?([A-Za-z0-9_]+)", 1, 0, false)
	return rs
}

type RuleSet struct {
	rules []*Rule
}

func NewRuleSet() *RuleSet {
	result := RuleSet {}
	result.Init()
	return &result
}

func (self *RuleSet) Init() {
	self.rules = rules()
}

func (self *RuleSet) CheckLine(tag *Tag, s string, lineCount, byteCount int) {
	for _, rule := range self.rules {
		applied := rule.Apply(tag, s, lineCount, byteCount)
		if applied { break }
	}
}
