package dictionary

import (
	"fmt"
	"strings"

	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"github.com/samber/lo"
)

type viewPortModel struct {
	ready    bool
	sentence string
	keys     addCommandKeyMap
	help     help.Model
	response scraper.Response
	viewport viewport.Model
}

var (
	borderStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
)

func MakeViewPortModel(response scraper.Response, sentence string) viewPortModel {
	return viewPortModel{
		keys:     keys,
		help:     help.New(),
		response: response,
		sentence: sentence,
	}
}

func (m viewPortModel) Init() tea.Cmd {
	return nil
}

func (m viewPortModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}
		// if key.Matches(msg, m.keys.Bottom) && m.ready {
		// 	m.viewport.GotoBottom()
		// }
		// if key.Matches(msg, m.keys.Top) && m.ready {
		// 	m.viewport.GotoTop()
		// }

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-2)
			m.viewport.SetContent(m.createContent())
			m.ready = true
		} else {
			m.viewport.Height = msg.Height
			m.viewport.Width = msg.Width
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m viewPortModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	helpView := m.help.View(m.keys)

	header := lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true).Render(fmt.Sprintf("The definition of the word %s", m.sentence))

	return fmt.Sprintf("%s\n%s\n%s", header, m.viewport.View(), helpView)
}

func (m viewPortModel) createContent() (content string) {
	for _, item := range m.response.Usages {
		content += m.separator(item) + "\n"

		for _, definition := range item.Definitions {
			sla := ""
			sla += "Definition: " + definition.WordDefinition

			if len(definition.Examples) > 0 {
				sla += "\n"
			}

			lo.ForEach(definition.Examples, func(example scraper.Example, _ int) {
				sla += "\n" + "Example: " + wordwrap.String(string(example), m.viewport.Width-9)
			})

			content += borderStyle.Render(sla) + "\n"
		}
	}

	return
}

func (m viewPortModel) separator(usage scraper.Usage) string {
	title := borderStyle.Render(usage.PartOfSpeech + " - " + usage.Language)
	line := strings.Repeat("-", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
