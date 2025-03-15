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

	assertNoError(t, err)

	assertEqualStrings(t, result, want)
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Got an unexpected error")
	}
}

func assertEqualStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
