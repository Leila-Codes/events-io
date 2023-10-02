package main

import (
	"github.com/Leila-Codes/events-io/sink"
	"github.com/Leila-Codes/events-io/source"
	"github.com/Leila-Codes/events-io/source/deserialize"
	"time"

	_ "github.com/lib/pq"
)

type ExampleJson struct {
	Timestamp     time.Time `json:"timestamp"`
	EventType     int       `json:"type"`
	EventCategory string    `json:"category"`
	EventData     string    `json:"data"`
}

func main() {

	input := source.NewKafkaDataSource[ExampleJson](
		[]string{"localhost:9092"},
		"test-topic-1",
		"testy-test-1",
		deserialize.Json[ExampleJson],
	)

	sink.SqlDataSink[ExampleJson](
		// channel input of events
		input,
		// sql.DB driver name
		"postgres",
		// driver specific connection string
		"postgres://postgres:postgres@localhost:5432/events_io_test?sslmode=disable",
		// insert single-row sql statement
		"INSERT INTO events (\"timestamp\", \"event_type\", \"data\") VALUES ($1, $2, $3)",
		// function that converts your struct into a series of values that match db columns
		func(e ExampleJson) []interface{} {
			return []interface{}{e.Timestamp, e.EventType, e.EventData}
		},
		// maximum number "per batch"
		50,
		// maximum time between batches
		30*time.Second,
	)

}
