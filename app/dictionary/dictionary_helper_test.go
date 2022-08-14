package dictionary

import (
	"reflect"
	"testing"
)

func TestCleanMeaningResults(t *testing.T) {
	md := func(definition, example string) Definition {
		return Definition{definition, example}
	}
	validateStructs := func(t testing.TB, got, want Meaning) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("meaning not clean, got %v want %v", got, want)
		}

	}

	t.Run("filter empty example", func(t *testing.T) {
		meaning := Meaning{PartOfSpeech: "verb", Definitions: []Definition{md("def1", "example"), md("def2", "")}}
		got := CleanMeaningResults(meaning)
		want := Meaning{PartOfSpeech: "verb", Definitions: []Definition{md("def1", "example")}}

		validateStructs(t, got, want)
	})

	t.Run("empty array", func(t *testing.T) {
		meaning := Meaning{PartOfSpeech: "noun", Definitions: []Definition{}}
		got := CleanMeaningResults(meaning)
		want := Meaning{PartOfSpeech: "noun", Definitions: []Definition{}}

		validateStructs(t, got, want)
	})

	t.Run("return empty definition array", func(t *testing.T) {
		meaning := Meaning{PartOfSpeech: "noun", Definitions: []Definition{md("", "example"), md("test", "")}}
		got := CleanMeaningResults(meaning)
		want := Meaning{PartOfSpeech: "noun", Definitions: []Definition{}}

		validateStructs(t, got, want)
	})
}
