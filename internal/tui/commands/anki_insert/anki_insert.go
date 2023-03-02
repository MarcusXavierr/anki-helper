package ankiinsert

import (
	"fmt"
	"strings"

	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

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

type model struct {
	ankiConfig ankiconnect.AnkiConfig
	list       list.Model
	isLoading  bool
	sentences  []ankiconnect.SentenceCards
	error      error
}

func MakeModel(config ankiconnect.AnkiConfig, sentences []ankiconnect.SentenceCards) model {
	item, sentences := pop(sentences)

	var err error

	if item.Error != nil {
		err = item.Error
	}

	if len(item.Cards) == 0 {
		err = fmt.Errorf("nao vai dar nao")
	}

	m := model{
		ankiConfig: config,
		list:       list.New(MakeCardItems(item), list.NewDefaultDelegate(), 0, 0),
		isLoading:  false,
		sentences:  sentences,
		error:      err,
	}
	m.list.Title = item.Sentence
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	if m.error != nil {
		return docStyle.Render("there was an error, press 'space' to continue")
	}

	return docStyle.Render(m.list.View())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(m.sentences) == 0 && m.error == nil {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.error != nil && msg.String() == tea.KeySpace.String() {
			h, w := m.list.Height(), m.list.Width()
			m := MakeModel(m.ankiConfig, m.sentences)
			m.list.SetSize(w, h)
			return m, nil
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd

}
