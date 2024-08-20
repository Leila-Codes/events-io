package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func kafkaReader[OUT interface{}](
	reader *kafka.Reader,
	output chan<- OUT,
	transformer KafkaTransformer[OUT],
) {
	for {
		m, err := reader.ReadMessage(context.TODO())

		if err != nil {
			log.Printf("Error reading message: %v", err)
		}

		output <- transformer(m)
	}
}

func DataSource[OUT interface{}](
	config kafka.ReaderConfig,
	bufferSize uint64,
	transformer KafkaTransformer[OUT],
) chan OUT {

	var (
		out    = make(chan OUT, bufferSize)
		reader = kafka.NewReader(config)
	)

	go kafkaReader(reader, out, transformer)

	return out
}
