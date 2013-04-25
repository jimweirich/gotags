package main

import (
	"fmt"
	"strings"
)

type Tag struct {
	path string
	data string
}

func NewTag(path string) *Tag {
	result := Tag { path: path }
	return &result
}

func (self *Tag) Add(tagname, line string, loc Location) {
	if tagname != "" {
		self.data = self.data + self.dataLineFor(tagname, line, loc)
	}
}

func(self *Tag) dataLineFor(tagname, line string, loc Location) string {
	line = self.firstLineOnly(line)
	result := fmt.Sprintf("%s\x7f%s\x01%d,%d\n", line, tagname, loc.LineCount, loc.ByteCount)
	return result
}

func (self *Tag) firstLineOnly(str string) string {
	splits := strings.Split(str, "\n")
	return strings.TrimRight(splits[0], "\n")
}

type tagWriter interface {
	WriteString(string) (int, error)
}

func (self *Tag) WriteOn(w tagWriter) {
	bytes := len(self.data)
	if bytes > 0 {
		w.WriteString("\x0c\n")
		w.WriteString(fmt.Sprintf("%s,%d\n", self.path, bytes))
		w.WriteString(self.data)
	}
}
