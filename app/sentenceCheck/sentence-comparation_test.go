package sentenceCheck

import (
	"testing"
)

func TestSentenceComparation(t *testing.T) {
	checkResultBool := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("Error, want %t but got %t", want, got)
		}
	}
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
