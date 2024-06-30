package kafka

import (
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

type MyEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Type      uint64    `json:"type"`
	Category  string    `json:"category"`
	Data      string    `json:"data"`
}

func TestKafkaDataSource(t *testing.T) {
	readerConfig := kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "event-io-test-1",
		Topic:       "test-topic-1",
		MinBytes:    1e3,
		MaxBytes:    1e6,
		StartOffset: kafka.FirstOffset,
	}

	messages := DataSource(
		readerConfig,
		10,
		RawMessages)

	events := DataSource(
		readerConfig,
		10,
		JsonValueDeserializer[MyEvent],
	)

	timestamps := DataSource(
		readerConfig,
		10,
		TimestampsOnly,
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
			t.Logf("Data Rx: %v", data)
			dataCount++
		case timestamp := <-timestamps:
			t.Logf("Timestamp Rx: %s", timestamp)
			timestampCount++
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
