package ccwc

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
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

func WordCount(filepath string, options WordCountOptions) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	counts := getCounts(string(content))

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

	wordCountString = fmt.Sprintf("%s %s", wordCountString, filepath)

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

func getCounts(content string) WordCountResults {
	return WordCountResults{
		bytes:      countBytes(content),
		characters: countCharacters(content),
		lines:      countLines(content),
		words:      countWords(content),
	}
}

func countBytes(content string) int {
	return len(content)
}

func countCharacters(content string) int {
	return utf8.RuneCountInString(content)
}

func countLines(content string) int {
	count := 0

	for _, c := range content {
		if c == '\n' {
			count++
		}
	}

	return count
}

func countWords(content string) int {
	count := 0
	inWord := false

	for _, c := range content {
		if unicode.IsSpace(c) {
			inWord = false
		} else {
			if !inWord {
				inWord = true
				count++
			}
		}
	}

	return count
}
