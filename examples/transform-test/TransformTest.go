package transform_test

import (
	"github.com/Leila-Codes/events-io/sink"
	"github.com/Leila-Codes/events-io/sink/serialize"
	"github.com/Leila-Codes/events-io/source"
	"github.com/Leila-Codes/events-io/source/deserialize"
	"github.com/Leila-Codes/events-io/transform"
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
	input := source.NewKafkaDataSource[ExampleJson](
		[]string{"localhost:9092"},
		"test-topic-1",
		"testy-reader-1",
		deserialize.Json[ExampleJson], // Deserializer for data values
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

	sink.NewKafkaDataSink[CategoryCount](
		middle,
		[]string{"localhost:9092"},
		"test-topic-2",
		serialize.Json[CategoryCount],
	)
}
