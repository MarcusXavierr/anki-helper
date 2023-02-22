package dictionary

import (
	"testing"

	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
)

func TestCreateConvertVisualization(t *testing.T) {

	examplePhrase := "Can you believe that scumbag Steve asked to sleep with her before even asking her name? Yeah, I'll traspass this line"
	exampleOne := []scraper.Example{scraper.Example(examplePhrase)}
	definitionTwo := scraper.Definition{WordDefinition: "A sleazy, disreputable or despicable person; lowlife.", Examples: exampleOne}
	response := scraper.Response{
		Usages: []scraper.Usage{
			{PartOfSpeech: "Noun", Language: "English", Definitions: []scraper.Definition{
				{WordDefinition: "A condom.", Examples: []scraper.Example{}},
				definitionTwo,
			}},
		},
	}

	initialState := viewPortModel{
		ready:    false,
		keys:     addCommandKeyMap{},
		help:     help.Model{},
		response: response,
		viewport: viewport.Model{Width: 100},
	}

	got := initialState.createContent()

	approvals.VerifyString(t, got)

}

// func TestUpdate(t *testing.T) {
// 	initialState := viewPortModel{
// 		ready:    false,
// 		keys:     addCommandKeyMap{},
// 		help:     help.Model{},
// 		response: scraper.Response{},
// 		viewport: viewport.Model{},
// 	}
//
// 	t.Run("goes to bottom when G is pressed", func(t *testing.T) {
// 		msg := tea.Key{Type: tea.KeyRunes, Runes: []rune{'G'}, Alt: false}
// 		newModel, cmd := initialState.Update(msg)
//
// 		model, err := newModel.(viewPortModel)
//
// 		if !err {
// 			t.Fatal(err)
// 		}
//
// 		got := model.viewport.AtBottom()
// 		want := true
//
// 		if got != want {
// 			t.Errorf("expected %t but got %t", want, got)
// 		}
//
// 		if cmd != nil {
// 			t.Errorf("expected no commands")
// 		}
// 	})
// }
