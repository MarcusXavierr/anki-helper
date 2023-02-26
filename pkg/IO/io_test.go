package IO

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"testing/fstest"
)

type fakeFile struct {
	FilePath string
}

func (f fakeFile) ReadFile() string {
	return "this is\njust\ntesting"
}

func TestPrinting(t *testing.T) {
	testFunction := func(t *testing.T, print func(io.Writer, string), color string) {

		buffer := &bytes.Buffer{}
		message := "testing message"
		print(buffer, message)

		got := buffer.String()
		want := string(color) + "testing message" + string(ColorReset)

		validateString(got, want, t)
	}

	t.Run("test red", func(t *testing.T) {
		testFunction(t, PrintRed, ColorRed)
	})

	t.Run("test gree", func(t *testing.T) {
		testFunction(t, PrintGreen, ColorGreen)
	})
}

func validateString(got string, want string, t testing.TB) {
	t.Helper()

	if got != want {
		t.Errorf("Error printing, got %q want %q", got, want)
	}
}

func TestReadFile(t *testing.T) {
	f := fakeFile{FilePath: "testing"}
	got := f.ReadFile()
	want := "this is\njust\ntesting"

	validateString(want, got, t)
}

func TestGetWords(t *testing.T) {
	fs := fstest.MapFS{
		"english_words.txt": {Data: []byte("hi\nmy\nname\nis\nmarcus")},
		"trash.txt":         {Data: []byte("hello\nworld\n")},
	}

	t.Run("read 3 lines from a file with 5 lines", func(t *testing.T) {
		got, _ := GetWords(fs, "english_words.txt", 3)
		want := []string{"hi", "my", "name"}

		assertArray(got, want, t)
	})

	t.Run("read 2 lines from a file with 3 lines", func(t *testing.T) {
		got, _ := GetWords(fs, "trash.txt", 5)
		want := []string{"hello", "world"}

		assertArray(got, want, t)
	})
}

func assertArray(got []string, want []string, t *testing.T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %q but got %q", want, got)
	}
}
