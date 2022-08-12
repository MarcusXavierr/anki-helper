package IO

import (
	"bytes"
	"io"
	"testing"
)

func TestPrinting(t *testing.T) {
	testFunction := func(t *testing.T, print func(io.Writer, string), color string) {

		buffer := &bytes.Buffer{}
		message := "testing message"
		print(buffer, message)

		got := buffer.String()
		want := string(color) + " testing message " + string(colorReset) + "\n"

		validateString(got, want, t)
	}

	t.Run("test red", func(t *testing.T) {
		testFunction(t, PrintRed, colorRed)
	})

	t.Run("test gree", func(t *testing.T) {
		testFunction(t, PrintGreen, colorGreen)
	})
}

func validateString(got string, want string, t testing.TB) {
	t.Helper()

	if got != want {
		t.Errorf("Error printing, got %q want %q", got, want)
	}
}
