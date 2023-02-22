package dictionary

import "github.com/charmbracelet/bubbles/key"

type addCommandKeyMap struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
}

func (k addCommandKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Quit}
}

func (k addCommandKeyMap) FullHelp() [][]key.Binding {
	return nil
}

var keys = addCommandKeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
