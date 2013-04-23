package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Tag struct {
	Path string
	Data string
}

func NewTag(path string) *Tag {
	result := Tag {}
	result.Init(path)
	return &result
}

func (self *Tag) Init(path string) {
	self.Path = path
	self.Data = ""
}

func (self *Tag) Add(tagname, line string, loc Location) {
	if tagname != "" {
		self.Data = self.Data + self.DataLineFor(tagname, line, loc)
	}
}

func(self *Tag) DataLineFor(tagname, line string, loc Location) string {
	line = strings.TrimRight(line, "\n")
	result := fmt.Sprintf("%s\x7f%s\x01%d,%d\n", line, tagname, loc.LineCount, loc.ByteCount)
	return result
}

func (self *Tag) WriteOn(w *bufio.Writer) {
	bytes := len(self.Data)
	if bytes > 0 {
		w.WriteString("\x0c\n")
		w.WriteString(fmt.Sprintf("%s,%d\n", self.Path, bytes))
		w.WriteString(self.Data)
	}
}
