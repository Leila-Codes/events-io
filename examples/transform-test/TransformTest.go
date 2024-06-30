package main

import (
	"github.com/Leila-Codes/events-io/plugins/kafka"
	"github.com/Leila-Codes/events-io/transform"
	kafka2 "github.com/segmentio/kafka-go"
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
	// input chan ExampleJSON
	input := kafka.DataSource[ExampleJson](
		kafka2.ReaderConfig{Brokers: []string{"localhost:9092"}, Topic: "test-topic-1", GroupID: "testy-reader-1"},
		1_000,
		kafka.JsonValueDeserializer[ExampleJson], // Deserializer for data values
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

	kafka.DataSink[CategoryCount](
		middle,
		&kafka2.Writer{
			Addr:  kafka2.TCP("localhost:9092"),
			Topic: "test-topic-2",
		},
		kafka.JsonValueSerializer[CategoryCount],
	)
}
