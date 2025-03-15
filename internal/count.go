package ccwc

import (
	"fmt"
	"os"
	"unicode"
)

type WordCountOptions struct {
	Characters bool
	Lines      bool
	Words      bool
}

type WordCountResults struct {
	characters int
	lines      int
	words      int
}

func WordCount(filepath string, options WordCountOptions) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	counts := getCounts(string(content))

	var wordCountString string

	if options.Characters {
		wordCountString = fmt.Sprintf("%d %s", counts.characters, filepath)
	}
	if options.Lines {
		wordCountString = fmt.Sprintf("%d %s", counts.lines, filepath)
	}
	if options.Words {
		wordCountString = fmt.Sprintf("%d %s", counts.words, filepath)
	}

	return wordCountString, nil
}

func getCounts(content string) WordCountResults {
	return WordCountResults{
		characters: countCharacters(content),
		lines:      countLines(content),
		words:      countWords(content),
	}
}

func countCharacters(content string) int {
	return len(content)
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
