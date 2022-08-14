package dictionary

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/MarcusXavierr/anki-helper/app/IO"
)

const cringeDefinition = "[{\"word\":\"cringe\",\"phonetic\":\"/kɹɪnd͡ʒ/\",\"phonetics\":[{\"text\":\"/kɹɪnd͡ʒ/\",\"audio\":\"https://api.dictionaryapi.dev/media/pronunciations/en/cringe-us.mp3\",\"sourceUrl\":\"https://commons.wikimedia.org/w/index.php?curid=5049283\",\"license\":{\"name\":\"BY-SA 3.0\",\"url\":\"https://creativecommons.org/licenses/by-sa/3.0\"}}],\"meanings\":[{\"partOfSpeech\":\"noun\",\"definitions\":[],\"synonyms\":[],\"antonyms\":[]},{\"partOfSpeech\":\"verb\",\"definitions\":[{\"definition\":\"To shrink, cower, tense or recoil, as in fear, disgust or embarrassment.\",\"synonyms\":[],\"antonyms\":[],\"example\":\"He cringed as the bird collided with the window.\"}],\"synonyms\":[],\"antonyms\":[]}],\"license\":{\"name\":\"CC BY-SA 3.0\",\"url\":\"https://creativecommons.org/licenses/by-sa/3.0\"},\"sourceUrls\":[\"https://en.wiktionary.org/wiki/cringe\"]}]"

func makeFakeServer(status int, response []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(response)
	}))

}
func TestGetDefinition(t *testing.T) {
	createResponse := func(defition, example, word string) DictionaryApiResponse {
		definition := Definition{
			Def:     defition,
			Example: example,
		}
		noumMeaning := Meaning{PartOfSpeech: "noun", Definitions: []Definition{}}
		verbMeaning := Meaning{PartOfSpeech: "verb", Definitions: []Definition{definition}}
		return DictionaryApiResponse{Word: word, Meanings: []Meaning{noumMeaning, verbMeaning}}
	}

	t.Run("Test cringe", func(t *testing.T) {
		server := makeFakeServer(200, []byte(cringeDefinition))
		got, _ := GetDefinition(server.URL)
		response := createResponse(
			"To shrink, cower, tense or recoil, as in fear, disgust or embarrassment.",
			"He cringed as the bird collided with the window.",
			"cringe",
		)
		want := response
		compareStructs(t, got, want)

	})

	t.Run("word dont exists", func(t *testing.T) {
		server := makeFakeServer(404, []byte("I dont know"))
		_, err := GetDefinition(server.URL)
		if err != IO.NotFoundError {
			t.Errorf("got %q want %q", err, IO.NotFoundError)
		}
	})
}

func compareStructs(t testing.TB, got, want DictionaryApiResponse) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
