package main

import (
	"github.com/Leila-Codes/events-io/sink"
	"time"

	_ "github.com/lib/pq"
)

type MyEvent struct {
	Username, Action, Message string
	Timestamp                 time.Time
}

func main() {
	eventStream := make(chan MyEvent)

	// generate random events onto eventStream
	go func() {
		for {
			eventStream <- MyEvent{
				Username: "leila-codes",
				Action:   "login",
				Message:  "User logged in",
			}

			// SLEEP
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// configure a SqlDataSink to dump these events
	sink.SqlDataSink(
		eventStream,
		"postgres",
		"postgres://postgres:postgres@localhost:5432/events_test?sslmode=disable",
		"INSERT INTO my_events (username, action, message) VALUES ($1, $2, $3)",
		func(event MyEvent) []interface{} {
			return []interface{}{event.Username, event.Action, event.Message}
		},
		10,
		3*time.Second,
	)
}
