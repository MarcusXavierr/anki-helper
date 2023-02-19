package IO

import (
	"bytes"
	"io"
	"testing"
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
