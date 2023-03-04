package ankiinsert

import (
	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
)

type model struct {
	ankiConfig     ankiconnect.AnkiConfig
	list           list.Model
	isLoading      bool
	sentences      []ankiconnect.SentenceCards
	actualSentence string
	error          error
	spinner        spinner.Model
	userFiles      IO.UserFilePath
}

func MakeModel(config ankiconnect.AnkiConfig, sentences []ankiconnect.SentenceCards, userFiles IO.UserFilePath) model {
	delegate := newItemDelegate(newDelegateKeyMap())
	item, sentences := pop(sentences)

	var err error

	if item.Error != nil {
		err = item.Error
	}

	if len(item.Cards) == 0 && err == nil {
		err = theresNoCards
	}

	m := model{
		ankiConfig:     config,
		list:           list.New(MakeCardItems(item), delegate, 0, 0),
		isLoading:      false,
		sentences:      sentences,
		actualSentence: item.Sentence,
		error:          err,
		spinner:        makeSpinner(),
		userFiles:      userFiles,
	}
	m.list.Title = item.Sentence
	m.list.SetSpinner(spinner.Monkey)
	return m
}
