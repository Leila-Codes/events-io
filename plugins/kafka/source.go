package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func kafkaReader[OUT interface{}](
	reader *kafka.Reader,
	output chan OUT,
	deserializer KafkaDeserializer[OUT],
) {
	for {
		m, err := reader.ReadMessage(context.TODO())
		if err != nil {
			log.Printf("Error reading message: %v", err)
		}

		output <- deserializer(m)
	}
}

func DataSource[OUT interface{}](
	config kafka.ReaderConfig,
	bufferSize uint64,
	deserializer KafkaDeserializer[OUT],
) chan OUT {

	var (
		out    = make(chan OUT, bufferSize)
		reader = kafka.NewReader(config)
	)

	go kafkaReader(reader, out, deserializer)

	return out
}
