package dictionary

import (
	"encoding/json"
	"errors"
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

func GetDefinition(url string) (DictionaryApiResponse, error) {
	return IO.HttpGetDictionary(parser, url)
}

func parser(input io.Reader, status int) (DictionaryApiResponse, error) {
	body, err := ioutil.ReadAll(input)
	if err != nil {
		return DictionaryApiResponse{}, err
	}

	var response []DictionaryApiResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return DictionaryApiResponse{}, err
	}

	if len(response) <= 0 {
		return DictionaryApiResponse{}, errors.New("No data fetched")
	}

	return response[0], nil

}
