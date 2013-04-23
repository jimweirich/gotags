package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Location struct {
	LineCount, ByteCount int
}

func processFile(writer *bufio.Writer, path string) {
	loc := Location { LineCount: 1, ByteCount: 0 }
	tag := NewTag(path)
	if filepath.Ext(path) != ".rb" {
		return
	}
	f, _ := os.Open(path)
	r := bufio.NewReader(f)
	var err error = nil
	var line string
	rset := NewRuleSet()
	for {
		line, err = r.ReadString('\n')
		if err != nil {
			break
		}
		rset.CheckLine(tag, line, loc)
		loc.LineCount++
		loc.ByteCount += len(line)
	}
	tag.WriteOn(writer)
}

func walkDir(writer *bufio.Writer, path string, info os.FileInfo, err error) error {
	if ! info.IsDir() {
		processFile(writer, path)
	}
	return nil
}

func main() {
	flag.Parse()

	fo, _ := os.Create("TAGS")
	defer fo.Close()

	writer := bufio.NewWriter(fo)
	defer writer.Flush()

	walkFunc := func(path string, info os.FileInfo, err error) error {
		return walkDir(writer, path, info, err)
	}

	var err error = nil

	if len(flag.Args()) == 0 {
		err = filepath.Walk(".", walkFunc)
	} else {
		for _, arg := range flag.Args() {
			err = filepath.Walk(arg, walkFunc)
		}
	}
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}
