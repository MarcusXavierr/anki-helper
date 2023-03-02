package ankiconnect

import (
	"github.com/MarcusXavierr/wiktionary-scraper/pkg/api"
	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
	"github.com/gofrs/uuid"
	"github.com/samber/lo"
)

type SentenceCards struct {
	Cards    []CardData
	Error    error
	Sentence string
}

func CreateCardDataMatrix(sentences []string, config AnkiConfig) []SentenceCards {
	resultChan := make(chan SentenceCards)

	cards := []SentenceCards{}
	for _, sent := range sentences {
		go func(sentence string) {
			response, err := api.GetDefinition("https://en.wiktionary.org/api/rest_v1/page/definition/", sentence)
			resultChan <- SentenceCards{
				Cards:    makeCardsFromResponse(response.Normalize(), sentence, config),
				Error:    err,
				Sentence: sentence,
			}
		}(sent)
	}

	for i := 0; i < len(sentences); i++ {
		r := <-resultChan
		if r.Error == nil {
			cards = append(cards, r)
		}
	}

	return cards
}

func makeCardsFromResponse(response scraper.Response, sentence string, config AnkiConfig) []CardData {
	normalizedResponse := response.Normalize()
	cards := lo.Map(normalizedResponse.Usages, func(item scraper.Usage, index int) []CardData {
		return mapDefinition(item.Definitions, sentence, config)
	})

	return flatten(cards)
}

func flatten[T any](lists [][]T) []T {
	var res []T
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func mapDefinition(definitions []scraper.Definition, sentence string, config AnkiConfig) []CardData {
	return lo.Map(definitions, func(item scraper.Definition, index int) CardData {
		id, _ := uuid.NewV4()
		return CardData{
			Sentence:   sentence,
			Example:    safeHead(item.Examples),
			Definition: item.WordDefinition,
			DeckName:   config.DeckName,
			ModelName:  config.ModelName,
			UUID:       id,
		}
	})
}

func safeHead(examples []scraper.Example) string {
	if len(examples) < 1 {
		return ""
	}

	return string(examples[0])
}
