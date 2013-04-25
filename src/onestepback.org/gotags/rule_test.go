package main

import (
	"onestepback.org/assert"
	"testing"
)

type fauxRecorder struct {
	name, def string
	loc Location
}
func (self *fauxRecorder) Add(tagname, defstring string, loc Location) {
	self.name = tagname
	self.def = defstring
	self.loc = loc
}

func TestRuleThatTriggers(t *testing.T) {
	r := NewRule("^[ \t]*def ([a-z]+)", 1, 0)
	recorder := &fauxRecorder { }

	r.Apply(recorder, "def foo(a, b)", Location { 3, 21 })

	assert.StringEqual(t, "foo", recorder.name)
	assert.StringEqual(t, "def foo", recorder.def)
	assert.IntEqual(t,  3, recorder.loc.LineCount)
	assert.IntEqual(t, 21, recorder.loc.ByteCount)
}

func TestRuleThatDoesNotTrigger(t *testing.T) {
	r := NewRule("^[ \t]*def ([a-z]+)", 1, 0)
	recorder := &fauxRecorder { }

	r.Apply(recorder, "dog = Dog.new", Location { 3, 21 })

	assert.StringEqual(t, "", recorder.name)
}

type fauxAdder struct {
	added bool
}
func (self *fauxAdder) Add(tag tagRecorder, tagname, defstring string, matches []string, loc Location) {
	self.added = true
}

func TestRuleWithAlternativeAddStrategy(t *testing.T) {
	adder := &fauxAdder { }
	r := NewRule("^[ \t]*def ([a-z]+)", 1, 0).With(adder)
	recorder := &fauxRecorder { }

	r.Apply(recorder, "def foo(a, b)", Location { 3, 21 })

	assert.True(t, adder.added)
}
