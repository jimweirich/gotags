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

func (self *Rule) Match(data string) (string, string, bool) {
	matches := self.Pattern.FindStringSubmatch(data)
	if len(matches) > 0 {
		return matches[self.TagIndex], matches[self.DefIndex], true
	}
	return "", "", false
}

func (self *Rule) Apply(tag *Tag, data string, loc Location) bool {
	tagname, defstring, ok := self.Match(data)
	if ok {
		if self.IsMulti {
			self.AddMulti(tag, tagname, defstring, loc)
		} else {
			self.AddSingle(tag, tagname, defstring, loc)
		}
		return true
	}
	return false
}

func (self *Rule) AddSingle(tag *Tag, tagname, defstring string, loc Location) {
	tag.Add(tagname, defstring, loc)
}

func (self *Rule) AddMulti(tag *Tag, tagname, defstring string, loc Location) {
	for _, name := range strings.Split(tagname, ",") {
		name = strings.Trim(name, " *:")
		tag.Add(name, defstring, loc)
	}
}

var RubyRules = []*Rule {
	NewRule("^[ \t]*(class|module)[ \t]+([A-Z][A-Za-z0-9_]+::)*([A-Z][A-Za-z0-9_]*)", 3, 0, false),
	NewRule("^[ \t]*def[ \t]+((self\\.)?[a-z0-9_]+(!?)?)", 1, 0, false),
	NewRule("^[ \t]*([A-Z][A-Za-z0-9_]*)[ \t]*=", 1, 0, false),
	NewRule("^[ \t]*attr_(reader|writer|accessor)[ \t]+([:a-z0-9_, ]+)", 2, 0, true),
	NewRule("^[ \t]*alias[ \t]+:?([A-Za-z0-9_]+)", 1, 0, false),
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
	self.rules = RubyRules
}

func (self *RuleSet) CheckLine(tag *Tag, s string, loc Location) {
	for _, rule := range self.rules {
		applied := rule.Apply(tag, s, loc)
		if applied { break }
	}
}
