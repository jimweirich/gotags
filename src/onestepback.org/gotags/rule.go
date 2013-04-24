package main

import (
	"strings"
	"regexp"
)

type TagAdder interface {
	Add(tag *Tag, tagname, defstring string, matches []string, loc Location)
}

type Rule struct {
	Pattern  *regexp.Regexp
	TagIndex int
	DefIndex int
	IsMulti  bool
	Adder    TagAdder
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
	self.Adder = AddSingleTag {}
	return self
}

func (self *Rule) With(adder TagAdder) *Rule {
	self.Adder = adder
	return self
}

func (self *Rule) Match(data string) (string, string, []string, bool) {
	matches := self.Pattern.FindStringSubmatch(data)
	if len(matches) > 0 {
		return matches[self.TagIndex], self.firstLineOnly(matches[self.DefIndex]), matches, true
	}
	return "", "", []string{}, false
}

func (self *Rule) firstLineOnly(str string) string {
	splits := strings.Split(str, "\n")
	return splits[0]
}

func (self *Rule) Apply(tag *Tag, data string, loc Location) bool {
	tagname, defstring, matches, ok := self.Match(data)
	if ok {
		self.Adder.Add(tag, tagname, defstring, matches, loc)
		return true
	}
	return false
}

// Basic add strategy used for 90% of the rules

type AddSingleTag struct {
}

func (self AddSingleTag) Add(tag *Tag, tagname, defstring string, matches []string, loc Location) {
	tag.Add(tagname, defstring, loc)
}
