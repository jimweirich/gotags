package main

import (
	"regexp"
)

type TagAdder interface {
	Add(tag tagRecorder, matches []string, loc Location)
}

type Rule struct {
	Pattern  *regexp.Regexp
	TagIndex int
	DefIndex int
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
	self.Adder = AddSingleTag {tagIndex, 0}
	return self
}

func (self *Rule) With(adder TagAdder) *Rule {
	self.Adder = adder
	return self
}

type tagRecorder interface {
	Add(tagname, defstring string, loc Location)
}

func (self *Rule) Apply(tag tagRecorder, data string, loc Location) bool {
	matches := self.Pattern.FindStringSubmatch(data)
	if len(matches) > 0 {
		self.Adder.Add(tag, matches, loc)
		return true
	}
	return false
}

// Basic add strategy used for most of the rules

type AddSingleTag struct {
	nameIndex, defIndex int
}

func (self AddSingleTag) Add(tag tagRecorder, matches []string, loc Location) {
	tag.Add(matches[self.nameIndex], matches[self.defIndex], loc)
}
