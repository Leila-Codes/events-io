package source

import (
	"context"
	"fmt"
	"github.com/Leila-Codes/events-io/source/deserialize"
	"github.com/segmentio/kafka-go"
	"os"
	"strings"
	"time"
)

func kafkaReader[OUT interface{}](config kafka.ReaderConfig, out chan OUT, deserializer deserialize.Deserializer[[]byte, OUT]) {
	reader := kafka.NewReader(config)

	for {
		ctx, cleanup := context.WithTimeout(context.Background(), 10*time.Second)
		m, err := reader.ReadMessage(ctx)
		cleanup()
		if err != nil {
			panic(err)
		}

		out <- deserializer(m.Value)
	}
}

func KafkaDataSource[OUT interface{}](config kafka.ReaderConfig, deserializer deserialize.Deserializer[[]byte, OUT]) chan OUT {
	out := make(chan OUT, 1_000)

	go kafkaReader[OUT](config, out, deserializer)

	return out
}

func DefaultConsumerFromEnv() (cfg kafka.ReaderConfig, err error) {

	var (
		bootstrapUrls, hasServer    = os.LookupEnv("BOOTSTRAP_SERVER")
		topicName, hasTopic         = os.LookupEnv("TOPIC_NAME")
		consumerId, hasGroup        = os.LookupEnv("CONSUMER_GROUP")
		startOffsetConfig, hasStart = os.LookupEnv("START_OFFSET")

		startOffset int64
	)

	if !hasServer {
		return cfg, fmt.Errorf("BOOTSTRAP_SERVER not set but is required in env")
	}

	if !hasTopic {
		return cfg, fmt.Errorf("TOPIC_NAME not set but is required in env")
	}

	if !hasGroup {
		return cfg, fmt.Errorf("CONSUMER_GROUP not set but is required in env")
	}

	if hasStart {
		switch strings.ToLower(startOffsetConfig) {
		case "earliest":
			startOffset = kafka.FirstOffset
		default:
			startOffset = kafka.LastOffset
		}
	}

	return kafka.ReaderConfig{
		Brokers:          strings.Split(bootstrapUrls, ","),
		GroupID:          consumerId,
		Topic:            topicName,
		MinBytes:         1e3, // 1KB
		MaxBytes:         1e6, // 1MB
		ReadBatchTimeout: 5 * time.Second,
		CommitInterval:   time.Second,
		StartOffset:      startOffset,
		ReadBackoffMax:   60 * time.Second,
		//Logger:                 nil,
		//ErrorLogger:            nil,
		MaxAttempts:           3,
		OffsetOutOfRangeError: false,
	}, nil
}
