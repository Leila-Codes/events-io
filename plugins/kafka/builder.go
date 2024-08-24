package kafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Builder[IN interface{}] func(IN) kafka.Message

func ToString[IN interface{}](input IN) kafka.Message {
	return kafka.Message{
		Value: []byte(fmt.Sprint(input)),
	}
}

func ToByte(input []byte) kafka.Message {
	return kafka.Message{
		Value: input,
	}
}

func ToKeyValue[IN interface{}](
	input IN,
	keyer func(IN) []byte,
	valuer func(IN) []byte,
) kafka.Message {
	return kafka.Message{
		Key:   keyer(input),
		Value: valuer(input),
	}
}
