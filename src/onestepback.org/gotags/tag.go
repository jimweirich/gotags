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

func (self *Tag) Add(tagname, line string, lineCount, byteCount int) {
	self.Data = self.Data + self.DataLineFor(tagname, line, lineCount, byteCount)
}

func(self *Tag) DataLineFor(tagname, line string, lineCount, byteCount int) string {
	line = strings.TrimRight(line, "\n")
	result := fmt.Sprintf("%s\x7f%s\x01%d,%d\n", line, tagname, lineCount, byteCount)
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
