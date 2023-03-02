package ankiinsert

import (
	"testing"

	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
)

func TestPopFunction(t *testing.T) {
	list := []ankiconnect.SentenceCards{
		{
			Cards:    []ankiconnect.CardData{},
			Error:    nil,
			Sentence: "test",
		},
		{
			Cards:    []ankiconnect.CardData{},
			Error:    nil,
			Sentence: "test 2",
		},
	}

	sentenceCard, newList := pop(list)
	if len(newList) != 1 {
		t.Errorf("wanted newList length to be %d, but got %d\n", 1, len(newList))
	}

	if sentenceCard.Sentence != "test 2" {
		t.Errorf("wanted sentence to be %q, but got %q\n", "test 2", sentenceCard.Sentence)
	}
}
