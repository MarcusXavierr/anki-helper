package dictionary

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/MarcusXavierr/anki-helper/app/IO"
)

type Response struct {
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

func GetDefinition(word, url string) (interface{}, error) {
	return IO.HttpGetDictionary(parser, url)
}

func parser(input io.Reader, status int) (interface{}, error) {
	body, err := ioutil.ReadAll(input)
	if err != nil {
		return Response{}, err
	}

	var response []Response
	json.Unmarshal(body, &response)
	return response, nil

}
