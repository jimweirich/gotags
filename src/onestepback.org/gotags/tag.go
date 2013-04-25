package main

import (
	"fmt"
	"strings"
)

type Tag struct {
	Path string
	Data string
}

func NewTag(path string) *Tag {
	result := Tag { Path: path, Data: "" }
	return &result
}

func (self *Tag) Add(tagname, line string, loc Location) {
	if tagname != "" {
		self.Data = self.Data + self.dataLineFor(tagname, line, loc)
	}
}

func(self *Tag) dataLineFor(tagname, line string, loc Location) string {
	line = strings.TrimRight(line, "\n")
	result := fmt.Sprintf("%s\x7f%s\x01%d,%d\n", line, tagname, loc.LineCount, loc.ByteCount)
	return result
}

type tagWriter interface {
	WriteString(string) (int, error)
}

func (self *Tag) WriteOn(w tagWriter) {
	bytes := len(self.Data)
	if bytes > 0 {
		w.WriteString("\x0c\n")
		w.WriteString(fmt.Sprintf("%s,%d\n", self.Path, bytes))
		w.WriteString(self.Data)
	}
}
