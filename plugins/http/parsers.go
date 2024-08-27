package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResponseParser[OUT interface{}] func(*http.Response) OUT

func RawBody(response *http.Response) ([]byte, error) {
	return io.ReadAll(response.Body)
}

func JsonBody[OUT interface{}](response *http.Response) (OUT, error) {
	output := new(OUT)
	err := json.NewDecoder(response.Body).Decode(output)
	return *output, err
}
