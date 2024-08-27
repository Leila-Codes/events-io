package http

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RequestBuilder[T interface{}] func(T) (*http.Request, error)

func RestRequestBuilder[IN interface{}](method, url string) RequestBuilder[IN] {
	return func(event IN) (*http.Request, error) {
		buff := &bytes.Buffer{}
		err := json.NewEncoder(buff).Encode(event)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest(method, url, buff)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		return req, err
	}
}
