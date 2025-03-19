package ccwc_test

import (
	"os"
	"testing"

	ccwc "github.com/eduardpeters/ccwc/internal"
)

var test_file_path = "test.txt"

func TestCountingFileBytesOutput(t *testing.T) {
	want := "342190 test.txt"
	test_options := ccwc.WordCountOptions{
		Bytes: true,
	}
	source, err := os.Open(test_file_path)
	assertNoError(t, err)

	result, err := ccwc.WordCount(source, test_file_path, test_options)

	assertNoError(t, err)

	assertEqualStrings(t, result, want)
}

func TestCountingFileCharactersOutput(t *testing.T) {
	want := "339292 test.txt"
	test_options := ccwc.WordCountOptions{
		Characters: true,
	}
	source, err := os.Open(test_file_path)
	assertNoError(t, err)

	result, err := ccwc.WordCount(source, test_file_path, test_options)

	assertNoError(t, err)

	assertEqualStrings(t, result, want)
}

func TestCountingFileLinesOutput(t *testing.T) {
	want := "7145 test.txt"
	test_options := ccwc.WordCountOptions{
		Lines: true,
	}
	source, err := os.Open(test_file_path)
	assertNoError(t, err)

	result, err := ccwc.WordCount(source, test_file_path, test_options)

	assertNoError(t, err)

	assertEqualStrings(t, result, want)
}

func TestCountingFileWordsOutput(t *testing.T) {
	want := "58164 test.txt"
	test_options := ccwc.WordCountOptions{
		Words: true,
	}
	source, err := os.Open(test_file_path)
	assertNoError(t, err)

	result, err := ccwc.WordCount(source, test_file_path, test_options)

	assertNoError(t, err)

	assertEqualStrings(t, result, want)
}

func TestCountingFileDefaultOutput(t *testing.T) {
	want := "  7145  58164 342190 test.txt"
	test_options := ccwc.WordCountOptions{
		Bytes:      false,
		Characters: false,
		Lines:      false,
		Words:      false,
	}
	source, err := os.Open(test_file_path)
	assertNoError(t, err)

	result, err := ccwc.WordCount(source, test_file_path, test_options)

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
