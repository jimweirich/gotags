package main

import (
)

type RuleSet struct {
	rules []*Rule
}

func (self *RuleSet) CheckLine(tag *Tag, s string, loc Location) {
	for _, rule := range self.rules {
		applied := rule.Apply(tag, s, loc)
		if applied { break }
	}
}
