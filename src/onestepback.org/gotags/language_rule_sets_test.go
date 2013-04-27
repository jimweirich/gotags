package main

import (
	"onestepback.org/assert"
	"testing"
)

func assertRecorded(t *testing.T, rec *fauxRecorder, methodName, defString string, lines, bytes int) {
	assert.StringEqual(t, methodName, rec.name)
	assert.StringEqual(t, defString, rec.def)
	assert.IntEqual(t, lines, rec.loc.LineCount)
	assert.IntEqual(t, bytes, rec.loc.ByteCount)
}

var RubyRules = &RuleSet { rules: RubyRulesList }

func TestRubyRuleWithClass(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "class Dog", Location { 12, 34 })

	assertRecorded(t, recorder, "Dog", "class Dog", 12, 34)
}

func TestRubyRuleWithScopedClass(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "class Scope::Dog", Location { 12, 34 })

	assertRecorded(t, recorder, "Dog", "class Scope::Dog", 12, 34)
}

func TestRubyRuleWithModule(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "module Mod", Location { 12, 34 })

	assertRecorded(t, recorder, "Mod", "module Mod", 12, 34)
}

func TestRubyRuleWithMethod(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "def foo()", Location { 12, 34 })

	assertRecorded(t, recorder, "foo", "def foo", 12, 34)
}

func TestRubyRuleWithClassMethod(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "def self.foo()", Location { 12, 34 })

	assertRecorded(t, recorder, "self.foo", "def self.foo", 12, 34)
	assert.IntEqual(t, 2, recorder.count)
	assert.StringEqual(t, "<foo><self.foo>", recorder.names)
}

func TestRubyRuleWithMethodAndQuestion(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "def foo?()", Location { 12, 34 })

	assertRecorded(t, recorder, "foo?", "def foo?", 12, 34)
}

func TestRubyRuleWithMethodAndBang(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "def foo!()", Location { 12, 34 })

	assertRecorded(t, recorder, "foo!", "def foo!", 12, 34)
}

func TestRubyRuleWithMethodAndEqual(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "def foo=(value)", Location { 12, 34 })

	assertRecorded(t, recorder, "foo=", "def foo=", 12, 34)
}

func TestRubyRuleWithAttrReader(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "attr_reader :a, :b", Location { 12, 34 })

	assertRecorded(t, recorder, "b", "attr_reader :a, :b", 12, 34)
	assert.IntEqual(t, 2, recorder.count)
	assert.StringEqual(t, "<a><b>", recorder.names)
}

func TestRubyRuleWithAttrWriter(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "attr_writer :a, :b", Location { 12, 34 })

	assertRecorded(t, recorder, "b", "attr_writer :a, :b", 12, 34)
	assert.IntEqual(t, 2, recorder.count)
	assert.StringEqual(t, "<a><b>", recorder.names)
}

func TestRubyRuleWithAttrAccessor(t *testing.T) {
	recorder := &fauxRecorder { }
	RubyRules.CheckLine(recorder, "attr_accessor :a, :b", Location { 12, 34 })

	assertRecorded(t, recorder, "b", "attr_accessor :a, :b", 12, 34)
	assert.IntEqual(t, 2, recorder.count)
	assert.StringEqual(t, "<a><b>", recorder.names)
}
