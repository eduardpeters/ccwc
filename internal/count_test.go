package ccwc_test

import (
	"testing"

	ccwc "github.com/eduardpeters/ccwc/internal"
)

func TestCharacterCount(t *testing.T) {
	t.Run("It counts characters in a string", func(t *testing.T) {
		testString := "Lorem ipsum"
		got := ccwc.GetCounts(testString)
		want := 11

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
