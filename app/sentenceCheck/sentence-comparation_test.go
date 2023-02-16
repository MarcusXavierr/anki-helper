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
	const want = false
	t.Run("Should return null if substring doesn't belongs to string", func(t *testing.T) {
		got := verifyIfSentenceExists("test", "testing\nI don't know what I'm doing")
		checkResult(t, got, want)
	})

	t.Run("should return false if string is empty", func(t *testing.T) {
		got := verifyIfSentenceExists("test", "")
		checkResult(t, got, want)
	})

	t.Run("Should return false if substring is empty", func(t *testing.T) {
		got := verifyIfSentenceExists("", "testing")
		checkResult(t, got, want)
	})
}

func TestTrueSentenceComparation(t *testing.T) {
	t.Run("Verifify with success if sentence exists on string", func(t *testing.T) {
		const want = true
		got := verifyIfSentenceExists("test", "test")
		checkResult(t, got, want)

		got = verifyIfSentenceExists("test", "testing\ntest\nnew test")
		checkResult(t, got, want)
	})

}

func TestCheckIfSentenceExists(t *testing.T) {
	buffer := &bytes.Buffer{}
	file, trashFile := fakeFile{}, fakeTrashFile{}

	t.Run("sentence exists", func(t *testing.T) {
		got := CheckIfSentenceExists(buffer, "my", file, trashFile)
		want := true
		checkResult(t, got, want)
	})

	t.Run("sentence dont exists", func(t *testing.T) {
		got := CheckIfSentenceExists(buffer, "testing", file, trashFile)
		want := false
		checkResult(t, got, want)
	})
}

func checkResult(t testing.TB, got bool, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Error, want %t but got %t", want, got)
	}
}
