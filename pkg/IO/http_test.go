package IO

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type fakeReponse struct {
	Status string
	Person fakePerson `json:"person"`
}

type fakePerson struct {
	Name string
	Age  int
}

var notFound = errors.New("Not found")

func makeFakeServer(status int, response []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(response)
	}))

}
func TestHttpGetDictionary(t *testing.T) {

	t.Run("got 200 code", func(t *testing.T) {
		fakeServer := makeFakeServer(200, []byte("{\"status\":\"200\",\"person\":{\"age\":18,\"name\":\"joaquim de lima\"}}"))
		fakePerson := fakePerson{Name: "joaquim de lima", Age: 18}

		got, _ := HttpGetDictionary(parser, fakeServer.URL)
		want := fakeReponse{Status: "200", Person: fakePerson}
		validateStruct(t, got, want)

	})

	t.Run("got an error", func(t *testing.T) {
		fakeServer := makeFakeServer(404, []byte("{ \"error\" }"))

		_, err := HttpGetDictionary(parser, fakeServer.URL)

		if err == nil {
			t.Errorf("should throw an error ")
		}
	})
}

func parser(input io.Reader, status int) (interface{}, error) {
	body, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}

	fmt.Println(body)
	var response fakeReponse
	json.Unmarshal(body, &response)
	return response, nil

}

func validateStruct(t testing.TB, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}

}
