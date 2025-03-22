package ccwc_test

import (
	"os"
	"strings"
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

func TestCountingReaderDefaultOutput(t *testing.T) {
	want := "  0  5 25"
	test_options := ccwc.WordCountOptions{
		Bytes:      false,
		Characters: false,
		Lines:      false,
		Words:      false,
	}
	source := strings.NewReader("this is some piped input?")

	result, err := ccwc.WordCount(source, "", test_options)

	assertNoError(t, err)

	assertEqualStrings(t, result, want)
}

func BenchmarkWordCount(b *testing.B) {
	benchmark_text := `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam posuere viverra ullamcorper. Pellentesque a mi sed leo porta maximus nec ut tortor. Lorem ipsum dolor sit amet, consectetur adipiscing elit. In convallis lacus sed aliquam maximus. Etiam interdum arcu in turpis volutpat, quis tincidunt orci iaculis. Ut sit amet eleifend tellus. Morbi condimentum blandit ipsum quis rhoncus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Vivamus lacinia leo vitae dapibus vulputate. Integer mauris orci, sollicitudin eu mollis id, sollicitudin eget felis. Duis sapien libero, dictum quis elementum eu, fringilla sit amet eros.

Etiam velit lacus, ultrices vitae lacus in, ornare pulvinar mi. Quisque fermentum malesuada lacinia. Ut semper arcu enim, ac tempor neque gravida a. Phasellus lacinia convallis nulla, non feugiat arcu dapibus a. Vivamus semper sollicitudin erat vitae aliquet. Aliquam imperdiet, neque quis elementum consequat, nulla mi tincidunt turpis, vitae tincidunt nibh felis a erat. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Sed convallis ligula sed ex ullamcorper tincidunt. Aliquam faucibus leo eget lorem semper, eget efficitur ligula rhoncus. Quisque a fermentum magna. Maecenas et erat non turpis accumsan pharetra. Nullam tristique tortor massa, non imperdiet velit tempor eu. Donec fringilla turpis eget est condimentum, sed tempus lectus ornare. Nulla a justo quam.

Nunc turpis nulla, facilisis sed auctor eget, consectetur quis felis. Maecenas eu sem ac mi rutrum ultrices nec et massa. Suspendisse eget felis at metus mollis pharetra vel non est. Morbi tempus lacus quis lacus vestibulum, sed dignissim enim maximus. Donec tincidunt pulvinar lorem a fermentum. Donec a dapibus magna. In lorem eros, euismod dignissim hendrerit sed, mollis vel eros. Phasellus vulputate risus eu tortor tincidunt facilisis.

Cras ac lacus at turpis vulputate auctor. Integer ornare maximus metus sed pharetra. In hendrerit faucibus lectus, eget blandit mauris pellentesque id. In vulputate sodales neque, non malesuada metus facilisis a. Integer vehicula vehicula ipsum, quis tempor erat hendrerit id. Vivamus hendrerit ligula molestie, lobortis libero sit amet, egestas lorem. Etiam quis justo consequat diam condimentum lobortis. Ut ipsum neque, interdum vulputate condimentum ut, tristique sit amet velit. Praesent a eleifend enim. Praesent molestie mi arcu, quis finibus augue congue a. Ut porta hendrerit orci, id dignissim ipsum fermentum et. Praesent egestas sagittis nulla.

Phasellus cursus sollicitudin mi eget malesuada. Aenean ac diam ac nulla congue mollis ac id elit. Nunc tellus magna, fermentum at consectetur ac, bibendum blandit ipsum. Nullam finibus tortor eget magna placerat porta. Sed finibus mauris enim, non vestibulum nibh hendrerit vitae. Sed quis dui eget augue porttitor imperdiet sed et magna. Ut sed sagittis ipsum. `
	benchmark_options := ccwc.WordCountOptions{
		Bytes:      true,
		Characters: true,
		Lines:      true,
		Words:      true,
	}

	for i := 0; i < b.N; i++ {
		source := strings.NewReader(benchmark_text)
		ccwc.WordCount(source, "", benchmark_options)
	}

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
