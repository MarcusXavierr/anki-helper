package sentenceCheck

import (
	"bytes"
	"testing"
)

type fakeTrashFile struct{}
type fakeFile struct{}

func (f fakeFile) ReadFile() string {
	return "my\nfake\nfile"
}

func (f fakeTrashFile) ReadFile() string {
	return "trash"
}

func TestSentenceComparation(t *testing.T) {
	t.Run("Verifify without success if sentence exists on a string", func(t *testing.T) {
		const want = false
		got := verifyIsSentenceExists("test", "testing\nI don't know what I'm doing")
		checkResultBool(t, got, want)

		got = verifyIsSentenceExists("test", "")
		checkResultBool(t, got, want)

		got = verifyIsSentenceExists("", "testing")
		checkResultBool(t, got, want)
	})

	t.Run("Verifify with success if sentence exists on string", func(t *testing.T) {
		const want = true
		got := verifyIsSentenceExists("test", "test")
		checkResultBool(t, got, want)

		got = verifyIsSentenceExists("test", "testing\ntest\nnew test")
		checkResultBool(t, got, want)
	})
}

func TestCheckIfSentenceExists(t *testing.T) {
	buffer := &bytes.Buffer{}
	file, trashFile := fakeFile{}, fakeTrashFile{}

	t.Run("sentence exists", func(t *testing.T) {
		got := CheckIfSentenceExists(buffer, "my", file, trashFile)
		want := true
		checkResultBool(t, got, want)
	})

	t.Run("sentence dont exists", func(t *testing.T) {
		got := CheckIfSentenceExists(buffer, "testing", file, trashFile)
		want := false
		checkResultBool(t, got, want)
	})
}

func checkResultBool(t testing.TB, got bool, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Error, want %t but got %t", want, got)
	}
}
