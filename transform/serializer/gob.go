package serializer

import (
	"bytes"
	"encoding/gob"

	"github.com/Leila-Codes/events-io/transform"
)

func gobSerializer[IN interface{}](input IN) []byte {
	buff := &bytes.Buffer{}
	dec := gob.NewEncoder(buff)
	err := dec.Encode(input)
	if err != nil {
		panic("GOB serializer error: " + err.Error())
	}

	return buff.Bytes()
}

// Gob serializer will attempt to serialize all events received from input chan
// and encode them into an in-memory bytes buffer before returning them on a new output channel.
func Gob[IN interface{}](input <-chan IN) chan []byte {
	return transform.Map(input, gobSerializer)
}
