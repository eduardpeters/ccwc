package ccwc

import (
	"fmt"
	"os"
)

type WordCountOptions struct {
	Characters bool
}

func WordCount(filepath string, options WordCountOptions) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	counts := GetCounts(string(content))

	wordCountString := fmt.Sprintf("%d %s", counts, filepath)

	return wordCountString, nil
}

func GetCounts(content string) int {
	return countCharacters(content)
}

func countCharacters(content string) int {
	return len(content)
}
