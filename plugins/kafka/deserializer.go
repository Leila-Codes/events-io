package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaTransformer[OUT interface{}] func(message kafka.Message) OUT

func RawMessage(message kafka.Message) kafka.Message {
	return message
}

func ByteValue(m kafka.Message) []byte {
	return m.Value
}

func StringValue(m kafka.Message) string {
	return string(m.Value)
}

func TimestampsOnly(m kafka.Message) time.Time {
	return m.Time
}
