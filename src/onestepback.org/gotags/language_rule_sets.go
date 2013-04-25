package main

import (
	"strings"
)

// Adding multilple, comma-separated names with the same location.

type AddMultipleTags struct {
	namesIndex, defIndex int
}

func (self AddMultipleTags) Add(tag tagRecorder, matches []string, loc Location) {
	names := matches[self.namesIndex]
	defstring := FirstLineOnly(matches[self.defIndex])
	for _, name := range strings.Split(names, ",") {
		name = strings.Trim(name, " \t\n*:")
		tag.Add(name, defstring, loc)
	}
}

// Adding a type-scoped go function

type AddGoClassTag struct {
	nameIndex, classIndex, defIndex int
}

func (self AddGoClassTag) Add(tag tagRecorder, matches []string, loc Location) {
	name := matches[self.nameIndex]
	classname := matches[self.classIndex]
	defstring := FirstLineOnly(matches[self.defIndex])
	tag.Add(name, defstring, loc)
	tag.Add(classname + "." + name, defstring, loc)
}

// Ruby Rules

var RubyRulesList = []*Rule {
	NewRule("^[ \t]*(class|module)[ \t]+([A-Z][A-Za-z0-9_]+::)*([A-Z][A-Za-z0-9_]*)", 3, 0),
	NewRule("^[ \t]*def[ \t]+((self\\.)?[A-Za-z0-9][A-Za-z0-9_]*(!?)?)", 1, 0),
	NewRule("^[ \t]*([A-Z][A-Za-z0-9_]*)[ \t]*=", 1, 0),
	NewRule("^[ \t]*attr_(reader|writer|accessor)[ \t]+([:A-Za-z0-9_, \t\n]+)", 2, 0).
		With(AddMultipleTags { 2, 0 }),
	NewRule("^[ \t]*alias(_method)?[ \t]+:?([A-Za-z0-9_]+)", 2, 0),
}

// Go Rules

var GoRulesList = []*Rule {
	NewRule("^(func|var)[ \t]+([A-Za-z0-9_]+)", 2, 0),
	NewRule("^type[ \t]+([A-Za-z0-9_]+)[ \t]+(struct|interface)", 1, 0),
	NewRule("^(func|var)[ \t]+\\([a-zA-Z0-9_]+[ \t*]+([A-Z][A-Za-z0-9_]+)\\)[ \t]*([A-Za-z0-9_]+)", 3, 0).
		With(AddGoClassTag {3, 2, 0}),
}

// Master list of all Rules

var Rules = map[string] *RuleSet {
	".rb": &RuleSet { rules: RubyRulesList },
	".go": &RuleSet { rules: GoRulesList },
}
