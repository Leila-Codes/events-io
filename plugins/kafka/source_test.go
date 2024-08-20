package kafka_test

import (
	"testing"
	"time"

	kafka "github.com/Leila-Codes/events-io/plugins/kafka"
	kafka2 "github.com/segmentio/kafka-go"
)

type MyEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Type      uint64    `json:"type"`
	Category  string    `json:"category"`
	Data      string    `json:"data"`
}

// TODO: Replace with mocks

// TestKafkaDataSource is a long running test that will simply log all messages in a variety of formats as received from a REAL kafka server.
// You will need:
//   - A kafka server running on localhost:9092
//   - A topic named "test-topic-1"
//   - A live stream of events (can be any textual data) streaming into this topic.
func TestKafkaDataSource(t *testing.T) {
	readerConfig := kafka2.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "event-io-test-1",
		Topic:       "test-topic-1",
		MinBytes:    1e3,
		MaxBytes:    1e6,
		StartOffset: kafka2.FirstOffset,
	}

	messages := kafka.DataSource(
		readerConfig,
		10,
		kafka.RawMessage,
	)

	events := kafka.DataSource(
		readerConfig,
		10,
		kafka.ByteValue,
	)

	timestamps := kafka.DataSource(
		readerConfig,
		10,
		kafka.TimestampsOnly,
	)

	var (
		rawCount       = 0
		timestampCount = 0
		dataCount      = 0
	)

	for rawCount < 10 && timestampCount < 10 && dataCount < 10 {
		select {
		case raw := <-messages:
			t.Logf("Raw Rx: %v", raw)
			rawCount++
		case data := <-events:
			t.Logf("Byte Rx: %v", data)
			dataCount++
		case timestamp := <-timestamps:
			t.Logf("Timestamp Rx: %s", timestamp)
			timestampCount++
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
