package main

import (
	"testing"
	"onestepback.org/assert"
)
import "strconv"

type StringIo struct {
	data string
}

func (self *StringIo) WriteString(s string) (int, error) {
	self.data += s
	return 0, nil
}

func TestTag(t *testing.T) {
	tag := NewTag("file.go")
	tag.Add("fun", "def fun", Location { 10, 123 })
	tag.Add("g", "def g", Location { 23, 150 })

	s := StringIo { "" }
	tag.WriteOn(&s)

	defstring :=
		"def fun\x7ffun\x0110,123\n" +
		"def g\x7fg\x0123,150\n"
	expected := "\x0c\n" +
		"file.go," + strconv.Itoa(len(defstring)) + "\n" +
		defstring

	assert.StringEqual(t, expected, s.data)
}
