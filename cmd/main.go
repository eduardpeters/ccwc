package main

import (
	"flag"
	"fmt"

	ccwc "github.com/eduardpeters/ccwc/internal"
)

func main() {
	characterCount := flag.Bool("c", false, "Count characters")
	flag.Parse()

	filePath := flag.Arg(0)

	options := ccwc.WordCountOptions{
		Characters: *characterCount,
	}

	result, err := ccwc.WordCount(filePath, options)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
