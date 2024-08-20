package deserializer

import (
	"encoding/json"

	"github.com/Leila-Codes/events-io/transform"
)

func jsonDeserialize[OUT interface{}](raw []byte) OUT {
	value := new(OUT)

	// attempt de-serialize, handle error (panic)
	err := json.Unmarshal(raw, value)
	if err != nil {
		panic(err)
	}

	// return output on output channel
	return *value
}

func Json[OUT interface{}](raw <-chan []byte) chan OUT {
	return transform.Map(raw, jsonDeserialize[OUT])
}
