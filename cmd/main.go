package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	ccwc "github.com/eduardpeters/ccwc/internal"
)

func main() {
	byteCount := flag.Bool("c", false, "Count bytes")
	characterCount := flag.Bool("m", false, "Count characters")
	lineCount := flag.Bool("l", false, "Count lines")
	wordCount := flag.Bool("w", false, "Count words")
	flag.Parse()

	filePath := flag.Arg(0)

	options := ccwc.WordCountOptions{
		Bytes:      *byteCount,
		Characters: *characterCount,
		Lines:      *lineCount,
		Words:      *wordCount,
	}

	var source io.Reader
	var err error
	if filePath != "" {
		source, err = os.Open(filePath)
	} else {
		source = os.Stdin
	}
	if err != nil {
		panic(err)
	}

	result, err := ccwc.WordCount(source, filePath, options)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
