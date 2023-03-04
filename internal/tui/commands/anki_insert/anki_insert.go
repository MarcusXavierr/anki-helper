package ankiinsert

import (
	"errors"
	"fmt"

	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var theresNoCards = errors.New("There's no Cards")

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) View() string {
	// tá feio, pode firar func sla
	if m.error == theresNoCards {
		m.list.Title += " - Could not find examples"
	} else if m.error == emptyListPop {
		return "Good bye!!\nPress q or space to escape\n"
	} else if m.error != nil {
		return docStyle.Render("there was an error, press 'space' to continue " + m.error.Error())
	}

	if m.isLoading {
		return fmt.Sprintf("\n\n %s loading... \n\n", m.spinner.View())
	}

	return docStyle.Render(m.list.View())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(m.sentences) == 0 && m.error != nil && m.error != theresNoCards {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == tea.KeySpace.String() {
			h, w := m.list.Height(), m.list.Width()
			m := MakeModel(m.ankiConfig, m.sentences, m.userFiles)
			m.list.SetSize(w, h)
			moveSentence(m)
			return m, nil
		}

		if msg.String() == tea.KeyCtrlC.String() {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case insertMsg:
		//Ok, como eu faço agr? Eu crio mais uma msg pra botar pra mimir (ou uso uma que tem), dps crio um comando que insere no anki, trata erro, move palavra de arquivo, e retorna uma msg falando oq deu (podendo escrever no error do model, sla). E por fim, mando um comando que finaliza o loading da lista
		// cmd := m.list.StartSpinner()
		return m, tea.Sequence(makeStartLoadMsg(), makeStopLoadMsg(msg.item))
	case startLoadMsg:
		m.isLoading = true
		return m, spinner.Tick

	case stopLoadMsg:
		// pode virar func
		err := msg.item.card.InsertCard("http://localhost:8765")
		if err == nil {
			err = IO.MoveSentenceToFile(m.userFiles.TrashFile, m.userFiles.WriteFile, msg.item.card.Sentence)
		}

		h, w := m.list.Height(), m.list.Width()
		m = MakeModel(m.ankiConfig, m.sentences, m.userFiles)
		m.list.SetSize(w, h)
		if m.error == nil {
			m.error = err
		}
		return m, nil
	}

	if m.isLoading {
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd

}
