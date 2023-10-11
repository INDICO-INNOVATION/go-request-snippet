package requestsnippet

import (
	"testing"
)

func TestRequest(t *testing.T) {
	request := &Request{
		Method:  "GET",
		URI:     "https://google.com",
		Body:    nil,
		Headers: nil,
	}

	response, err := request.Call()
	if err != nil {
		t.Fatalf("Error on request")
	}
	t.Log(response)
}
