package ccwc

import (
	"fmt"
	"os"
)

type WordCountOptions struct {
	Characters bool
}

type WordCountResults struct {
	characters int
	lines      int
}

func WordCount(filepath string, options WordCountOptions) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	counts := getCounts(string(content))

	wordCountString := fmt.Sprintf("%d %s", counts.characters, filepath)

	return wordCountString, nil
}

func getCounts(content string) WordCountResults {
	return WordCountResults{
		characters: countCharacters(content),
		lines:      countLines(content),
	}
}

func countCharacters(content string) int {
	return len(content)
}

func countLines(content string) int {
	if len(content) == 0 {
		return 0
	}

	count := 1

	for _, c := range content {
		if c == '\n' {
			count++
		}
	}

	return count
}
