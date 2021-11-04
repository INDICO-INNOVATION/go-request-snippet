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

type StandardLogging struct {
	User       string      `json:"token"`
	Message    string      `json:"message"`
	Reason     string      `json:"reason"`
	RemoteIp   string      `json:"ipaddress"`
	Method     string      `json:"method"`
	Urlpath    string      `json:"route"`
	StatusCode int         `json:"status_code"`
	Payload    interface{} `json:"payload"`
}

type GenericGcp struct {
	Message   string `json:"message"`
	ProjectId string `json:"project_id"`
	TopicId   string `json:"topic_id"`
}

type TotpValidation struct {
	Valid bool
}

type GenericSfResponse struct {
	Header   []Header
	Response interface{}
}

func Request(method string, endpoint string, body io.Reader, headers []Header) (*GenericResponse, error) {

	request, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		// Log(&StandardLogging{
		// 	Message:    "Error on Request (1)",
		// 	StatusCode: http.StatusInternalServerError,
		// 	Reason:     err.Error(),
		// 	Payload:    request,
		// }, "critical")
		return nil, err
	}

	for _, slice := range headers {
		request.Header.Set(slice.Key, slice.Value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		// Log(&StandardLogging{
		// 	Message:    "Error on Request (2)",
		// 	StatusCode: http.StatusInternalServerError,
		// 	Reason:     err.Error(),
		// 	Payload:    response,
		// }, "critical")
		return nil, err
	}

	defer response.Body.Close()

	responsePayload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// Log(&StandardLogging{
		// 	Message:    "Error on Request (3)",
		// 	StatusCode: http.StatusInternalServerError,
		// 	Reason:     err.Error(),
		// 	Payload:    responsePayload,
		// }, "critical")
		return nil, err
	}

	return &GenericResponse{Message: responsePayload, StatusCode: response.StatusCode}, nil
}
