package main

import (
	"github.com/Leila-Codes/events-io/plugins/kafka"
	"github.com/Leila-Codes/events-io/plugins/sql_io"
	kafka2 "github.com/segmentio/kafka-go"
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

	input := kafka.DataSource[ExampleJson](
		kafka2.ReaderConfig{
			Topic:   "test-topic-1",
			GroupID: "testy-test-1",
			Brokers: []string{"localhost:9092"},
		},
		1_000,
		kafka.JsonValueDeserializer[ExampleJson],
	)

	sql_io.DataSink[ExampleJson](
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
