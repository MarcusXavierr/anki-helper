package IO

import (
	"errors"
	"io"
	"net/http"
)

var NotFoundError = errors.New("Info not found while fetching data")

func HttpGetDictionary(parser func(io.Reader, int) (interface{}, error), url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return "", NotFoundError
	}

	return parser(resp.Body, resp.StatusCode)
}
