package dictionary

func CleanMeaningResults(meaning Meaning) Meaning {
	definitions := []Definition{}
	definitions = removeEmptyDefinitions(meaning, definitions)

	return Meaning{PartOfSpeech: meaning.PartOfSpeech, Definitions: definitions}
}

func (d DictionaryApiResponse) Normalize() DictionaryApiResponse {
	meanings := []Meaning{}
	for _, meaning := range d.Meanings {
		meaning = CleanMeaningResults(meaning)
		if len(meaning.Definitions) > 0 {
			meanings = append(meanings, meaning)
		}
	}

	return DictionaryApiResponse{Word: d.Word, Meanings: meanings}
}

func removeEmptyDefinitions(meaning Meaning, definitions []Definition) []Definition {
	for _, definition := range meaning.Definitions {
		if definition.Def != "" && definition.Example != "" {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}
