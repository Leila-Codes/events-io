package source

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func kafkaReader[OUT interface{}](config kafka.ReaderConfig, out chan OUT, deserializer ValueDeserializer[OUT]) {
	reader := kafka.NewReader(config)

	for {
		ctx, cleanup := context.WithTimeout(context.Background(), 10*time.Second)
		m, err := reader.ReadMessage(ctx)
		cleanup()
		if err != nil {
			panic(err)
		}

		out <- deserializer(m)
	}
}

func KafkaDataSource[OUT interface{}](config kafka.ReaderConfig, deserializer ValueDeserializer[OUT]) chan OUT {
	out := make(chan OUT, 1_000)

	go kafkaReader[OUT](config, out, deserializer)

	return out
}
