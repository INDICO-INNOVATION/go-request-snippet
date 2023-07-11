package requestsnippet

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"
	"strings"
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

type Request struct {
	Method  string
	URI     string
	Body    io.Reader
	Headers []Header
	SkipTLS bool
}

func (req *Request) Call() (*GenericResponse, error) {
	request, err := http.NewRequest(strings.ToUpper(req.Method), req.URI, req.Body)
	if err != nil {
		return nil, err
	}

	for _, slice := range req.Headers {
		request.Header.Set(slice.Key, slice.Value)
	}

	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: req.SkipTLS,
			},
		},
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responsePayload, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &GenericResponse{Message: responsePayload, StatusCode: response.StatusCode}, nil
}
