package main

import (
	"strings"
)

// Adding multilple, comma-separated names with the same location.

type AddMultipleTags struct {
}

func (self AddMultipleTags) Add(tag *Tag, tagname, defstring string, matches []string, loc Location) {
	for _, name := range strings.Split(tagname, ",") {
		name = strings.Trim(name, " \t\n*:")
		tag.Add(name, defstring, loc)
	}
}

// Adding a type-scoped go function

type AddGoClassTag struct {
	classIndex int
}

func (self AddGoClassTag) Add(tag *Tag, tagname, defstring string, matches []string, loc Location) {
	tag.Add(tagname, defstring, loc)
	tag.Add(matches[self.classIndex] + "." + tagname, defstring, loc)
}

// Ruby Rules

var RubyRulesList = []*Rule {
	NewRule("^[ \t]*(class|module)[ \t]+([A-Z][A-Za-z0-9_]+::)*([A-Z][A-Za-z0-9_]*)", 3, 0),
	NewRule("^[ \t]*def[ \t]+((self\\.)?[A-Za-z0-9][A-Za-z0-9_]*(!?)?)", 1, 0),
	NewRule("^[ \t]*([A-Z][A-Za-z0-9_]*)[ \t]*=", 1, 0),
	NewRule("^[ \t]*attr_(reader|writer|accessor)[ \t]+([:A-Za-z0-9_, \t\n]+)", 2, 0).
		With(AddMultipleTags { }),
	NewRule("^[ \t]*alias(_method)?[ \t]+:?([A-Za-z0-9_]+)", 2, 0),
}

// Go Rules

var GoRulesList = []*Rule {
	NewRule("^(func|var)[ \t]+([A-Za-z0-9_]+)", 2, 0),
	NewRule("^type[ \t]+([A-Za-z0-9_]+)[ \t]+(struct|interface)", 1, 0),
	NewRule("^(func|var)[ \t]+\\([a-zA-Z0-9_]+[ \t*]+([A-Z][A-Za-z0-9_]+)\\)[ \t]*([A-Za-z0-9_]+)", 3, 0).
		With(AddGoClassTag {classIndex: 2}),
}

// Master list of all Rules

var Rules = map[string] *RuleSet {
	".rb": NewRuleSet(RubyRulesList),
	".go": NewRuleSet(GoRulesList),
}
