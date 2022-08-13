package dictionary

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/MarcusXavierr/anki-helper/app/IO"
)

type DictionaryApiResponse struct {
	Word     string
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	PartOfSpeech string
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Def     string `json:"definition"`
	Example string
}

func GetDefinition(word, url string) ([]DictionaryApiResponse, error) {
	return IO.HttpGetDictionary(parser, url)
}

func parser(input io.Reader, status int) ([]DictionaryApiResponse, error) {
	body, err := ioutil.ReadAll(input)
	if err != nil {
		return []DictionaryApiResponse{}, err
	}

	var response []DictionaryApiResponse
	json.Unmarshal(body, &response)
	return response, nil

}
