package deserializer

import (
	"bytes"
	"encoding/gob"

	"github.com/Leila-Codes/events-io/transform"
)

func gobDeserializer[OUT interface{}](input []byte) OUT {
	out := new(OUT)
	buff := bytes.NewReader(input)

	dec := gob.NewDecoder(buff)
	err := dec.Decode(out)
	if err != nil {
		panic("GOB deserializer error: " + err.Error())
	}

	return *out
}

// Gob deserializer implementation attempts to decode every event received in bytes
// and deserializes using encoding/gob to deserialize into a new instance of OUT type.
func Gob[OUT interface{}](input <-chan []byte) chan OUT {
	return transform.Map(input, gobDeserializer[OUT])
}
