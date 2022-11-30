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
		want := Meaning{PartOfSpeech: "noun", Definitions: []Definition{md("test", "")}}

		validateStructs(t, got, want)
	})
}

func TestNormalize(t *testing.T) {
	md := func(definition, example string) Definition {
		return Definition{definition, example}
	}

	t.Run("normal test", func(t *testing.T) {
		response := makeResponse([]Meaning{
			makeMeaning([]Definition{md("", "")}, "verb"),
			makeMeaning([]Definition{md("test", "testing")}, "noun"),
		}, "word")

		got := response.Normalize()

		want := DictionaryApiResponse{Word: "word", Meanings: []Meaning{
			Meaning{PartOfSpeech: "noun", Definitions: []Definition{
				Definition{Def: "test", Example: "testing"}},
			},
		}}

		compareStructs(t, got, want)
	})

	t.Run("should take first definition of a partOfSpeech if all examples are null", func(t *testing.T) {
		response := makeResponse([]Meaning{
			makeMeaning([]Definition{md("to test something", ""), md("test", "")}, "noun"),
			makeMeaning([]Definition{md("just test", ""), md("test", "")}, "verb"),
		}, "testing")

		got := response.Normalize()
		want := DictionaryApiResponse{Word: "testing", Meanings: []Meaning{
			Meaning{PartOfSpeech: "noun", Definitions: []Definition{
				Definition{Def: "to test something", Example: ""},
			}},
			{PartOfSpeech: "verb", Definitions: []Definition{
				{Def: "just test", Example: ""},
			}},
		}}

		compareStructs(t, got, want)
	})
}
func makeResponse(meanings []Meaning, word string) DictionaryApiResponse {
	return DictionaryApiResponse{Word: word, Meanings: meanings}
}

func makeMeaning(definitions []Definition, partOfSpeech string) Meaning {
	verbMeaning := Meaning{PartOfSpeech: partOfSpeech, Definitions: definitions}
	return verbMeaning
}
