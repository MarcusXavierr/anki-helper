package dictionary

func CleanMeaningResults(meaning Meaning) Meaning {
	definitions := []Definition{}
	definitions = removeEmptyDefinitions(meaning, definitions)

	return Meaning{PartOfSpeech: meaning.PartOfSpeech, Definitions: definitions}
}

func removeEmptyDefinitions(meaning Meaning, definitions []Definition) []Definition {
	for _, definition := range meaning.Definitions {
		if definition.Def != "" && definition.Example != "" {
			definitions = append(definitions, definition)
		}
	}
	return definitions
}
