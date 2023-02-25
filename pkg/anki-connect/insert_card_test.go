package ankiconnect

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gofrs/uuid"
)

func TestInsertCard(t *testing.T) {
	card := CardData{Sentence: "test", Example: "example"}
	t.Run("status ok provides no response", func(t *testing.T) {
		server := makeFakeServer(http.StatusOK, []byte(successfulResponse))
		err := card.InsertCard(server.URL)

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should return error", func(t *testing.T) {
		server := makeFakeServer(http.StatusOK, []byte(errorResponse))
		err := card.InsertCard(server.URL)

		if err == nil {
			t.Fatal("should return an error")
		}
	})
}

func TestCreateStruct(t *testing.T) {
	randomId, _ := uuid.NewV4()

	card := CardData{Sentence: "mamaco", Example: "mamaco", Definition: "A cooler macaco", UUID: randomId}
	got, err := makePayload(card)

	if err != nil {
		t.Fatal(err)
	}
	want := addNoteRequest{
		Action:  "addNote",
		Version: 6,
		Params: addNoteParams{
			Note: note{
				DeckName:  "Padrão",
				ModelName: "Basic",
				Fields: noteFields{
					Front: "mamaco",
					Back:  "A cooler macaco",
				},
				Audio: []noteAudio{
					{
						URL:      "https://translate.google.com/translate_tts?ie=UTF-8&tl=en&q=mamaco&client=tw-ob",
						Filename: fmt.Sprintf("mamaco_%s.mp3", randomId),
						Fields:   []string{"Front"},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}
}

const successfulResponse = `{"result": 1677353601247, "error": null}`
const errorResponse = `{"result": null, "error": "'deckname'"}`

func makeFakeServer(status int, response []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(response)
	}))

}
