package main

import (
	"github.com/Leila-Codes/events-io/sink"
	"github.com/Leila-Codes/events-io/source"
	"github.com/Leila-Codes/events-io/transform"
	"github.com/segmentio/kafka-go"
	"time"
)

type CategoryCount struct {
	Category string `json:"category"`
	Count    int64  `json:"count"`
}

type ExampleJson struct {
	Timestamp     time.Time `json:"timestamp"`
	EventType     int       `json:"type"`
	EventCategory string    `json:"category"`
	EventData     string    `json:"data"`
}

func main() {
	// Begin asynchronous streaming of data from Kafka
	input := source.KafkaDataSource[ExampleJson](
		kafka.ReaderConfig{
			Brokers:          []string{"localhost:9092"},
			GroupID:          "testy-reader-1",
			GroupTopics:      nil,
			Topic:            "test-topic-1",
			MinBytes:         1e2,
			MaxBytes:         1e6,
			StartOffset:      kafka.LastOffset,
			ReadBatchTimeout: time.Second,
			CommitInterval:   time.Second,
		}, // Kafka Consumer Configuration
		source.JsonDeserializer[ExampleJson], // Deserializer for data values
	)

	categoryCounts := map[string]int64{}

	middle := transform.Map[ExampleJson, CategoryCount](
		input,
		func(record ExampleJson) CategoryCount {
			if _, exists := categoryCounts[record.EventCategory]; !exists {
				categoryCounts[record.EventCategory] = 0
			}

			categoryCounts[record.EventCategory]++

			return CategoryCount{
				Category: record.EventCategory,
				Count:    categoryCounts[record.EventCategory],
			}
		})

	sink.KafkaDataSink[CategoryCount](
		kafka.WriterConfig{
			Brokers:       []string{"localhost:9092"},
			Topic:         "test-topic-2",
			MaxAttempts:   3,
			QueueCapacity: 0,
			BatchSize:     0,
			BatchBytes:    1e3,
			BatchTimeout:  5 * time.Second,
			RequiredAcks:  1,
			Async:         true,
		},
		middle,
		sink.JsonSerializer[CategoryCount])
}
