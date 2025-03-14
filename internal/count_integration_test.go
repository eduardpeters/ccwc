package ccwc_test

import (
	"testing"

	ccwc "github.com/eduardpeters/ccwc/internal"
)

var test_file_path = "test.txt"

func TestCountingFileCharactersOutput(t *testing.T) {
	want := "342190 test.txt"
	test_options := ccwc.WordCountOptions{
		Characters: true,
	}

	result, err := ccwc.WordCount(test_file_path, test_options)

	if err != nil {
		t.Fatal("Got an unexpected error")
	}

	if result != want {
		t.Errorf("got %q, want %q", result, want)
	}
}
