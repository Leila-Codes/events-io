package kafka

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KafkaDeserializer[OUT interface{}] func(message kafka.Message) OUT

func RawMessages(message kafka.Message) kafka.Message {
	return message
}

func ByteValueDeserializer(m kafka.Message) []byte {
	return m.Value
}

func StringValueDeserializer(m kafka.Message) string {
	return string(m.Value)
}

func JsonValueDeserializer[OUT interface{}](m kafka.Message) OUT {
	var (
		out = new(OUT)
		err = json.Unmarshal(m.Value, out)
	)

	if err != nil {
		log.Fatal("Failed to deserialize kafka event: ", err)
	}

	return *out
}

func TimestampsOnly(m kafka.Message) time.Time {
	return m.Time
}
