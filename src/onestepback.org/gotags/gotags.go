package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func processFile(writer *bufio.Writer, path string) {
	tag := NewTag(path)
	ext := filepath.Ext(path)
	if ext == ".rake" || filepath.Base(path) == "Rakefile"  {
		ext = ".rb"
	}
	rset := Rules[ext]
	if rset != nil {
		source, err := OpenLineSource(path)
		if err != nil {
			fmt.Println("Error opening file '" + path + "': " + err.Error())
			return
		}
		defer source.Close()
		for {
			line, err := source.ReadLine()
			if err != nil {
				break
			}
			rset.CheckLine(tag, line, source.Loc)
		}
		tag.WriteOn(writer)
	}
}

func walkDir(writer *bufio.Writer, path string, info os.FileInfo, err error) error {
	if info != nil && ! info.IsDir() {
		processFile(writer, path)
	}
	return nil
}

var version = "1.1.0"

func main() {
	var showVersion bool = false
	var showHelp    bool = false

	flag.BoolVar(&showVersion, "v",    false, "Display the version number")
	flag.BoolVar(&showHelp,    "h",    false, "Display help text")
	flag.BoolVar(&showHelp,    "help", false, "Display help text")
	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if showHelp {
		fmt.Println("Usage: gotags [options] [file...]")
		fmt.Println("")
		fmt.Println("Options are:")
		flag.PrintDefaults()
		os.Exit(0)
	}

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
		os.Exit(-1)
	}
}
