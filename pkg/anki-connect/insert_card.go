package ankiconnect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
)

type CardData struct {
	Sentence, Example, Definition, DeckName, ModelName string
	UUID                                               uuid.UUID
}

func (c CardData) InsertCard(url string) error {
	requestBody, err := json.Marshal(makePayload(c))

	if err != nil {
		return err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	defer response.Body.Close()

	return validateReponse(response)
}

func validateReponse(response *http.Response) error {
	if response.StatusCode != 200 {
		return fmt.Errorf("Could not get response, status code: %d", response.StatusCode)
	}

	var result map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return err
	}

	if errorValue, ok := result["error"].(string); ok && errorValue != "" {
		return fmt.Errorf("Error: %s", errorValue)
	}

	return nil
}
