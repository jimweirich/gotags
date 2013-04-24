package main

import (
	"strings"
)


type MultiAddStrategy struct {
}

func (self MultiAddStrategy) Add(tag *Tag, tagname, defstring string, matches []string, loc Location) {
	for _, name := range strings.Split(tagname, ",") {
		name = strings.Trim(name, " \t\n*:")
		tag.Add(name, defstring, loc)
	}
}

type GoClassAddStrategy struct {
	classIndex int
}

func (self GoClassAddStrategy) Add(tag *Tag, tagname, defstring string, matches []string, loc Location) {
	tag.Add(tagname, defstring, loc)
	tag.Add(matches[self.classIndex] + "." + tagname, defstring, loc)
}

var RubyRulesList = []*Rule {
	NewRule("^[ \t]*(class|module)[ \t]+([A-Z][A-Za-z0-9_]+::)*([A-Z][A-Za-z0-9_]*)", 3, 0),
	NewRule("^[ \t]*def[ \t]+((self\\.)?[A-Za-z0-9][A-Za-z0-9_]*(!?)?)", 1, 0),
	NewRule("^[ \t]*([A-Z][A-Za-z0-9_]*)[ \t]*=", 1, 0),
	NewRule("^[ \t]*attr_(reader|writer|accessor)[ \t]+([:A-Za-z0-9_, \t\n]+)", 2, 0).
		With(MultiAddStrategy { }),
	NewRule("^[ \t]*alias(_method)?[ \t]+:?([A-Za-z0-9_]+)", 2, 0),
}

var GoRulesList = []*Rule {
	NewRule("^(func|var)[ \t]+([A-Za-z0-9_]+)", 2, 0),
	NewRule("^type[ \t]+([A-Za-z0-9_]+)[ \t]+(struct|interface)", 1, 0),
	NewRule("^(func|var)[ \t]+\\([a-zA-Z0-9_]+[ \t*]+([A-Z][A-Za-z0-9_]+)\\)[ \t]*([A-Za-z0-9_]+)", 3, 0).
		With(GoClassAddStrategy {classIndex: 2}),
}

var Rules = map[string] *RuleSet {
	".rb": NewRuleSet(RubyRulesList),
	".go": NewRuleSet(GoRulesList),
}
