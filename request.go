package requestsnippet

import (
	"io"
	"io/ioutil"
	"net/http"
)

type GenericResponse struct {
	Message    []byte `json:"message"`
	StatusCode int    `json:"status_code"`
}

type ErrorResponse struct {
	Message     string `json:"error"`
	Description string `json:"description"`
	StatusCode  int    `json:"status_code"`
}

type Header struct {
	Key   string
	Value string
}

func Request(method string, endpoint string, body io.Reader, headers []Header) (*GenericResponse, error) {

	request, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	for _, slice := range headers {
		request.Header.Set(slice.Key, slice.Value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responsePayload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &GenericResponse{Message: responsePayload, StatusCode: response.StatusCode}, nil
}
