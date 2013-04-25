package main

import (
	"strings"
	"regexp"
)

type TagAdder interface {
	Add(tag tagRecorder, tagname, defstring string, matches []string, loc Location)
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
	result.init(pattern, tagIndex, defIndex)
	return &result
}

func (self *Rule) init(pattern string, tagIndex, defIndex int) *Rule {
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

func (self *Rule) match(data string) (string, string, []string, bool) {
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

type tagRecorder interface {
	Add(tagname, defstring string, loc Location)
}

func (self *Rule) Apply(tag tagRecorder, data string, loc Location) bool {
	tagname, defstring, matches, ok := self.match(data)
	if ok {
		self.Adder.Add(tag, tagname, defstring, matches, loc)
	}
	return ok
}

// Basic add strategy used for most of the rules

type AddSingleTag struct {
}

func (self AddSingleTag) Add(tag tagRecorder, tagname, defstring string, matches []string, loc Location) {
	tag.Add(tagname, defstring, loc)
}
