package source

import (
	"fmt"
	"github.com/Leila-Codes/events-io/source/deserialize"
	"github.com/segmentio/kafka-go"
	"os"
	"testing"
	"time"
)

type ExampleJson struct {
	Timestamp time.Time `json:"timestamp"`
	EventType int       `json:"type"`
	EventData string    `json:"data"`
}

var kafkaDefaults = kafka.ReaderConfig{
	Brokers:  []string{"localhost:9092"},
	GroupID:  "testy-consumer-group",
	Topic:    "test-topic-1",
	MinBytes: 1e3,
	MaxBytes: 1e6,
}

func TestKafkaDataSource_String(t *testing.T) {
	data := KafkaDataSource[string](
		kafkaDefaults,
		deserialize.String)

	for i := 0; i < 10; i++ {
		fmt.Println(<-data)
	}
}

func TestJsonDeserializer_Json(t *testing.T) {
	data := KafkaDataSource[ExampleJson](
		kafkaDefaults,
		deserialize.Json[ExampleJson],
	)

	for i := 0; i < 10; i++ {
		fmt.Println(<-data)
	}
}

func TestDefaultConsumerFromEnv(t *testing.T) {
	os.Setenv("BOOTSTRAP_SERVER", "localhost:9092")
	os.Setenv("TOPIC_NAME", "test-topic-1")
	os.Setenv("CONSUMER_GROUP", "test-stream-reader")
	os.Setenv("START_OFFSET", "earliest")

	config, err := DefaultConsumerFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", config)
}
