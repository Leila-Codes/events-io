package main

import (
	"time"

	"github.com/Leila-Codes/events-io/plugins/kafka"
	"github.com/Leila-Codes/events-io/transform"
	"github.com/Leila-Codes/events-io/transform/deserializer"
	"github.com/Leila-Codes/events-io/transform/serializer"
	"github.com/Leila-Codes/events-io/util"
	kafka2 "github.com/segmentio/kafka-go"
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
	// Begin asynchronous streaming of raw event data from Kafka (message.Value []byte)
	raw, err := kafka.DataSource(
		kafka2.ReaderConfig{Brokers: []string{"localhost:9092"}, Topic: "test-topic-1", GroupID: "testy-reader-1"},
		1_000,
		kafka.ByteValue, // Deserializer for data values
	)

	go util.PanicHandler(err)

	// input chan ExampleJSON
	input := deserializer.Json[ExampleJson](raw)

	// convert single events into count events for a particular category.
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

	output := serializer.Json(middle)

	// write count events back to another kafka topic.
	kafka.DataSink(
		output,
		&kafka2.Writer{
			Addr:  kafka2.TCP("localhost:9092"),
			Topic: "test-topic-2",
		},
		kafka.ToByte,
	)
}
