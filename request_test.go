package requestsnippet

import (
	"testing"
)

func TestRequest(t *testing.T) {
	response, err := Request("GET", "https://google.com", nil, nil)
	if err != nil {
		t.Fatalf("Error on request")
	}
	t.Log(response)
}
