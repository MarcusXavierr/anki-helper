package IO

import (
	"errors"
	"io"
	"net/http"
)

var NotFoundError = errors.New("Info not found while fetching data")

func HttpGetDictionary[T any](parser func(io.Reader, int) (T, error), url string) (T, error) {
	var noop T
	resp, err := http.Get(url)
	if err != nil {
		return noop, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return noop, NotFoundError
	}

	return parser(resp.Body, resp.StatusCode)
}
