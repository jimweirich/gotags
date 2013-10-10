package main

import (
	"bufio"
	"os"
)

type LineSource struct {
	File *os.File
	BufferedReader *bufio.Reader
	Loc Location
	NextLoc Location
}

func OpenLineSource(path string) (*LineSource, error) {
	result := LineSource { }
	err := result.Init(path)
	return &result, err
}

func (self *LineSource) Init(path string) error {
	var err error
	self.File, err = os.Open(path)
	if err == nil {
		self.BufferedReader = bufio.NewReader(self.File)
	}
	self.Loc = NewLocation()
	self.NextLoc = NewLocation()
	return err
}

func (self *LineSource) ReadLine() (string, error) {
	self.Loc = self.NextLoc
	line := ""
	var err error
	var data string
	for {
		data, err = self.BufferedReader.ReadString('\n')
		if err != nil { break }
		self.NextLoc = self.NextLoc.Bump(data)
		line = line + data
		if ! self.endsInComma(line) { break }
	}
	return line, err
}

func (self *LineSource) Close() {
	self.File.Close()
}

func (self *LineSource) endsInComma(line string) bool {
	i := len(line)-1
	for i >= 0 {
		if line[i] == ',' {
			return true
		}
		if line[i] != '\n' && line[i] != ' ' {
			return false
		}
		i--
	}
	return false
}
