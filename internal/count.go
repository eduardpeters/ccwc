package ccwc

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type WordCountOptions struct {
	Bytes      bool
	Characters bool
	Lines      bool
	Words      bool
}

type WordCountResults struct {
	bytes      int
	characters int
	lines      int
	words      int
}

const (
	DOUBLE_SPACING = "  "
	SINGLE_SPACING = " "
)

func WordCount(source io.Reader, filepath string, options WordCountOptions) (string, error) {
	counts, err := getCounts(*bufio.NewReader(source))
	if err != nil {
		return "", err
	}

	optionCount := countOptions(options)
	if optionCount == 0 {
		options.Lines = true
		options.Words = true
		options.Bytes = true

		optionCount = 3
	}

	wordCountString := ""
	spacing := ""
	if options.Lines {
		if optionCount > 1 {
			spacing = DOUBLE_SPACING
		}
		wordCountString = fmt.Sprintf("%s%s%d", wordCountString, spacing, counts.lines)
	}
	if options.Words {
		if optionCount > 1 {
			spacing = DOUBLE_SPACING
		}
		wordCountString = fmt.Sprintf("%s%s%d", wordCountString, spacing, counts.words)
	}
	if options.Characters {
		if optionCount > 1 {
			spacing = SINGLE_SPACING
		}
		wordCountString = fmt.Sprintf("%s%s%d", wordCountString, spacing, counts.characters)
	}
	if options.Bytes {
		if optionCount > 1 {
			spacing = SINGLE_SPACING
		}
		wordCountString = fmt.Sprintf("%s%s%d", wordCountString, spacing, counts.bytes)
	}

	if filepath != "" {
		wordCountString = fmt.Sprintf("%s %s", wordCountString, filepath)
	}

	return wordCountString, nil
}

func countOptions(options WordCountOptions) int {
	optionCount := 0
	if options.Bytes {
		optionCount++
	}
	if options.Characters {
		optionCount++
	}
	if options.Lines {
		optionCount++
	}
	if options.Words {
		optionCount++
	}

	return optionCount
}

func getCounts(reader bufio.Reader) (WordCountResults, error) {
	results := WordCountResults{}

	for {
		line, err := reader.ReadString('\n')
		lineLength := len(line)
		if lineLength == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return results, err
		}

		results.bytes += lineLength

		wordCount := 0
		inWord := false

		for _, c := range line {
			results.characters++

			if unicode.IsSpace(c) {
				if c == '\n' {
					results.lines++
				}
				inWord = false
			} else {
				if !inWord {
					inWord = true
					wordCount++
				}
			}
		}

		results.words += wordCount

		if err != nil {
			if err == io.EOF {
				break
			}
			return results, err
		}
	}

	return results, nil
}
