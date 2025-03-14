package main

import (
	"fmt"
	"os"

	ccwc "github.com/eduardpeters/ccwc/internal"
)

func main() {
	filePath := os.Args[2]

	options := ccwc.WordCountOptions{
		Characters: true,
	}

	result, err := ccwc.WordCount(filePath, options)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
