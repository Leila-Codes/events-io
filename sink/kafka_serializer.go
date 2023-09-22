package sink

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type ValueSerializer[IN interface{}] func(in IN) kafka.Message

func JsonSerializer[IN interface{}](in IN) kafka.Message {
	data, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	return kafka.Message{
		Key:   nil,
		Value: data,
	}
}

type Stringable interface {
	String() string
}

func StringSerializer[IN Stringable](in IN) kafka.Message {
	return kafka.Message{
		Key:   nil,
		Value: []byte(in.String()),
	}
}
