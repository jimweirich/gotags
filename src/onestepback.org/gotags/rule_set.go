package main

import (
	"strings"
	"regexp"
)

type Strategy interface {
	Add(tag *Tag, tagname, defstring string, loc Location)
}

type Rule struct {
	Pattern  *regexp.Regexp
	TagIndex int
	DefIndex int
	IsMulti  bool
	AddStrategy Strategy
}

func NewRule(pattern string, tagIndex, defIndex int) *Rule {
	result := Rule {}
	result.Init(pattern, tagIndex, defIndex)
	return &result
}

func (self *Rule) Init(pattern string, tagIndex, defIndex int) *Rule {
	self.Pattern = regexp.MustCompile(pattern)
	self.TagIndex = tagIndex
	self.DefIndex = defIndex
	self.AddStrategy = SingleAddStrategy {}
	return self
}

func (self *Rule) With(s Strategy) *Rule {
	self.AddStrategy = s
	return self
}

func (self *Rule) Match(data string) (string, string, bool) {
	matches := self.Pattern.FindStringSubmatch(data)
	if len(matches) > 0 {
		return matches[self.TagIndex], self.firstLineOnly(matches[self.DefIndex]), true
	}
	return "", "", false
}

func (self *Rule) firstLineOnly(str string) string {
	splits := strings.Split(str, "\n")
	return splits[0]
}

func (self *Rule) Apply(tag *Tag, data string, loc Location) bool {
	tagname, defstring, ok := self.Match(data)
	if ok {
		self.AddStrategy.Add(tag, tagname, defstring, loc)
		return true
	}
	return false
}

type SingleAddStrategy struct {
}

func (self SingleAddStrategy) Add(tag *Tag, tagname, defstring string, loc Location) {
	tag.Add(tagname, defstring, loc)
}

type MultiAddStrategy struct {
}

func (self MultiAddStrategy) Add(tag *Tag, tagname, defstring string, loc Location) {
	for _, name := range strings.Split(tagname, ",") {
		name = strings.Trim(name, " \t\n*:")
		tag.Add(name, defstring, loc)
	}
}

type ClassAddStrategy struct {
}

func (self ClassAddStrategy) Add(tag *Tag, tagname, defstring string, loc Location) {
	tag.Add(tagname, defstring, loc)
}

type RuleSet struct {
	rules []*Rule
}

func NewRuleSet(listOfRules []*Rule) *RuleSet {
	result := RuleSet { }
	result.Init(listOfRules)
	return &result
}

func (self *RuleSet) Init(rules []*Rule) {
	self.rules = rules
}

func (self *RuleSet) CheckLine(tag *Tag, s string, loc Location) {
	for _, rule := range self.rules {
		applied := rule.Apply(tag, s, loc)
		if applied { break }
	}
}

var RubyRulesList = []*Rule {
	NewRule("^[ \t]*(class|module)[ \t]+([A-Z][A-Za-z0-9_]+::)*([A-Z][A-Za-z0-9_]*)", 3, 0),
	NewRule("^[ \t]*def[ \t]+((self\\.)?[A-Za-z0-9][A-Za-z0-9_]*(!?)?)", 1, 0),
	NewRule("^[ \t]*([A-Z][A-Za-z0-9_]*)[ \t]*=", 1, 0),
	NewRule("^[ \t]*attr_(reader|writer|accessor)[ \t]+([:A-Za-z0-9_, \t\n]+)", 2, 0).With(MultiAddStrategy{}),
	NewRule("^[ \t]*alias(_method)?[ \t]+:?([A-Za-z0-9_]+)", 2, 0),
}

var GoRulesList = []*Rule {
	NewRule("^(func|var)[ \t]+([A-Za-z0-9_]+)", 2, 0),
	NewRule("^type[ \t]+([A-Za-z0-9_]+)[ \t]+(struct|interface)", 1, 0),
	NewRule("^(func|var)[ \t]+\\([a-zA-Z0-9_]+[ \t*]+([A-Z][A-Za-z0-9_]+)\\)[ \t]*([A-Za-z0-9_]+)", 3, 0).
		With(ClassAddStrategy {}),
}

var Rules = map[string] *RuleSet {
	".rb": NewRuleSet(RubyRulesList),
	".go": NewRuleSet(GoRulesList),
}
