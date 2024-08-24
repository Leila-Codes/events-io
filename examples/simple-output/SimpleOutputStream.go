package main

import (
	"math/rand"
	"time"

	"github.com/Leila-Codes/events-io/plugins/kafka"
	"github.com/Leila-Codes/events-io/transform/serializer"
	kafka2 "github.com/segmentio/kafka-go"
)

type ExampleJson struct {
	Timestamp     time.Time `json:"timestamp"`
	EventType     int       `json:"type"`
	EventCategory string    `json:"category"`
	EventData     string    `json:"data"`
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ=@!"

func randString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	test := make(chan ExampleJson, 100)

	data := serializer.Json(test)

	go kafka.DataSink(
		data,
		&kafka2.Writer{
			Addr:  kafka2.TCP("localhost:9092"),
			Topic: "test-stream-1",
		},
		kafka.ToByte,
	)

	// Generate some test messages
	for {
		test <- ExampleJson{
			Timestamp:     time.Now(),
			EventType:     rand.Int(),
			EventData:     randString(rand.Intn(16) + 4),
			EventCategory: string(letters[rand.Intn(len(letters))]),
		}
		time.Sleep(time.Millisecond)
	}
}
