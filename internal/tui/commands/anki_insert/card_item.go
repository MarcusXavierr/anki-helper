package ankiinsert

import (
	"fmt"
	"strings"

	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
	"github.com/charmbracelet/bubbles/list"
)

type cardItem struct {
	card ankiconnect.CardData
}

func MakeCardItems(sentenceCards ankiconnect.SentenceCards) []list.Item {
	listItems := []list.Item{}
	for _, card := range sentenceCards.Cards {
		listItems = append(listItems, cardItem{card: card})
	}

	return listItems
}

func (c cardItem) Title() string {
	return fmt.Sprintf("%q", strings.TrimSpace(c.card.Example))
}

func (c cardItem) Description() string {
	return c.card.Definition
}

func (c cardItem) FilterValue() string {
	return c.card.Example
}
