package source

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type ValueDeserializer[OUT interface{}] func(m kafka.Message) OUT

func JsonDeserializer[OUT interface{}](m kafka.Message) OUT {
	v := new(OUT)
	err := json.Unmarshal(m.Value, v)
	if err != nil {
		panic(err)
	}

	return *v
}
