package main

import (
	"flag"
	"fmt"

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

	result, err := ccwc.WordCount(filePath, options)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
