package serializer

import (
	"encoding/json"

	"github.com/Leila-Codes/events-io/transform"
)

func jsonSerialize[IN interface{}](event IN) []byte {
	data, err := json.Marshal(event)

	if err != nil {
		panic(err)
	}

	return data
}

func Json[IN interface{}](input <-chan IN) chan []byte {
	return transform.Map(input, jsonSerialize)
}
