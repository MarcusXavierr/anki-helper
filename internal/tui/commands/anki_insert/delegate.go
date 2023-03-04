package ankiinsert

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func insert(item cardItem) func() tea.Msg {
	return func() tea.Msg {
		return insertMsg{item}
	}
}

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = makeUpdateFunc(keys)

	help := []key.Binding{keys.choose, keys.next}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

func makeUpdateFunc(keys *delegateKeyMap) func(msg tea.Msg, m *list.Model) tea.Cmd {
	return func(msg tea.Msg, m *list.Model) tea.Cmd {
		item, ok := m.SelectedItem().(cardItem)

		if !ok {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			if key.Matches(msg, keys.choose) {
				return insert(item)
			}
		}

		return nil
	}
}

type delegateKeyMap struct {
	choose key.Binding
	next   key.Binding
}

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.next,
	}
}

func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.next,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		next: key.NewBinding(
			key.WithKeys("space"),
			key.WithHelp("space", "goto next sentence"),
		),
	}
}
