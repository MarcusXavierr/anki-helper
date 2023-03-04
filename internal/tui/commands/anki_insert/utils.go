package ankiinsert

import (
	"fmt"

	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
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

func moveSentence(model model) error {
	if model.error != theresNoCards {
		return nil
	}

	return IO.MoveSentenceToFile(model.userFiles.ManualInsertFile, model.userFiles.WriteFile, model.actualSentence)
}

func makeSpinner() spinner.Model {
	spin := spinner.New()
	spin.Spinner = spinner.Monkey
	spin.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return spin
}
