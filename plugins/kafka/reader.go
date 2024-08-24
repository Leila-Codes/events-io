package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type Reader[OUT interface{}] func(message kafka.Message) OUT

func RawMessage(m kafka.Message) kafka.Message {
	return m
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
