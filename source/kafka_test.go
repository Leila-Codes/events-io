package source

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

type ExampleJson struct {
	Timestamp time.Time `json:"timestamp"`
	EventType int       `json:"type"`
	EventData string    `json:"data"`
}

func TestJsonDeserializer(t *testing.T) {
	data := KafkaDataSource[ExampleJson](
		kafka.ReaderConfig{
			Brokers:  []string{"localhost:9092"},
			GroupID:  "testy-consumer-group",
			Topic:    "test-topic-1",
			MinBytes: 1e3,
			MaxBytes: 1e6,
		},
		JsonDeserializer[ExampleJson],
	)

	for i := 0; i < 10; i++ {
		fmt.Println(<-data)
	}
}
