package requestsnippet

import (
	"testing"
)

func TestRequest(t *testing.T) {
	request := &Request{
		Method:  "GET",
		URI:     "https://dog.ceo/api/breeds/image/random",
		Body:    nil,
		Headers: nil,
		SkipTLS: true,
	}

	response, err := request.Call()
	if err != nil {
		t.Fatalf("Error on request: %s", err.Error())
	}

	t.Log(string(response.Message))
}
