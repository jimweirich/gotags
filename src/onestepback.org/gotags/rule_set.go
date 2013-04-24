package main

import (
)

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
