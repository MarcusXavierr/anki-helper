package ankiinsert

import (
	"fmt"

	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
)

var emptyListPop = fmt.Errorf("tried to pop an empty list")

func pop(slice []ankiconnect.SentenceCards) (ankiconnect.SentenceCards, []ankiconnect.SentenceCards) {
	if len(slice) == 0 {
		return ankiconnect.SentenceCards{
			Cards:    []ankiconnect.CardData{},
			Error:    emptyListPop,
			Sentence: "",
		}, slice
	}
	return slice[len(slice)-1], slice[:len(slice)-1]
}
