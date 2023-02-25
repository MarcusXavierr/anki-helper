package ankiconnect

import "fmt"

func makePayload(c CardData) addNoteRequest {
	url := fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&tl=en&q=%s&client=tw-ob", c.Example)
	return addNoteRequest{
		Action:  "addNote",
		Version: 6,
		Params: addNoteParams{
			Note: note{
				DeckName:  c.DeckName,
				ModelName: c.ModelName,
				Fields: noteFields{
					Front: "mamaco",
					Back:  "A cooler macaco",
				},
				Audio: []noteAudio{
					{
						URL:      url,
						Filename: fmt.Sprintf("%s_%s.mp3", c.Sentence, c.UUID),
						Fields:   []string{"Front"},
					},
				},
			},
		},
	}

}

type noteFields struct {
	Front string `json:"Front"`
	Back  string `json:"Back"`
}

type noteAudio struct {
	URL      string   `json:"url"`
	Filename string   `json:"filename"`
	Fields   []string `json:"fields"`
}

type note struct {
	DeckName  string      `json:"deckName"`
	ModelName string      `json:"modelName"`
	Fields    noteFields  `json:"fields"`
	Audio     []noteAudio `json:"audio"`
}

type addNoteParams struct {
	Note note `json:"note"`
}

type addNoteRequest struct {
	Action  string        `json:"action"`
	Version int           `json:"version"`
	Params  addNoteParams `json:"params"`
}
