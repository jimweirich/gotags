package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func processFile(writer *bufio.Writer, path string) {
	lineCount := 1
	byteCount := 0
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
		rset.CheckLine(tag, line, lineCount, byteCount)
		lineCount++
		byteCount += len(line)
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
	var concurrent int
	flag.IntVar   (&concurrent, "c", 20,    "Number of concurrent fetchers")

	flag.Parse()

	fo, _ := os.Create("TAGS")
	defer fo.Close()

    writer := bufio.NewWriter(fo)
	defer writer.Flush()

	walkFunc := func(path string, info os.FileInfo, err error) error {
		return walkDir(writer, path, info, err)
	}

	for _, arg := range flag.Args() {
		err := filepath.Walk(arg, walkFunc)
		if err != nil {
			fmt.Println("Error while processing " + arg + ", Error: " + err.Error())
		}
	}
}
