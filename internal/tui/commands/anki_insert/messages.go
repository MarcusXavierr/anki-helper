package ankiinsert

import tea "github.com/charmbracelet/bubbletea"

type insertMsg struct {
	item cardItem
}
type startLoadMsg struct{}

type stopLoadMsg struct {
	item cardItem
}

func makeStartLoadMsg() func() tea.Msg {
	return func() tea.Msg { return startLoadMsg{} }
}

func makeStopLoadMsg(item cardItem) func() tea.Msg {
	return func() tea.Msg { return stopLoadMsg{item} }
}
