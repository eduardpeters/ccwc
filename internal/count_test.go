package ccwc

import (
	"testing"
)

func TestCharacterCount(t *testing.T) {
	t.Run("It counts characters in a string", func(t *testing.T) {
		testString := "Lorem ipsum"
		got := getCounts(testString)
		want := 11

		assertEqualCounts(t, got.characters, want)
	})
}

func TestLineCount(t *testing.T) {
	var lineTests = []struct {
		name  string
		input string
		want  int
	}{
		{"Should count no lines", "", 0},
		{"Should still count no lines", "Lorem ipsum", 0},
		{"Should count three lines", "Lorem\nipsum\ndolor sit amet", 2},
		{"Should count one line", "Lorem\nipsum", 1},
		{"Should count two lines", "\n\n", 2},
	}

	for _, tt := range lineTests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCounts(tt.input)

			assertEqualCounts(t, got.lines, tt.want)
		})
	}
}

func assertEqualCounts(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
