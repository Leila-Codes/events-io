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

var (
	defaultConsumerConfig = kafka.ReaderConfig{
		QueueCapacity:          1_000, // 1K messages
		MinBytes:               1e3,   // 1KB
		MaxBytes:               1e6,   // 1MB
		ReadBatchTimeout:       5 * time.Second,
		ReadLagInterval:        30 * time.Second,
		CommitInterval:         time.Second,
		PartitionWatchInterval: 30 * time.Second,
		WatchPartitionChanges:  true,
		ReadBackoffMin:         time.Second,
		ReadBackoffMax:         time.Minute,
		//Logger:                 nil,
		//ErrorLogger:            nil,
		//IsolationLevel:         0,
		MaxAttempts:           5,
		OffsetOutOfRangeError: true,
	}
)

func kafkaReader[OUT interface{}](config kafka.ReaderConfig, out chan OUT, deserializer deserialize.Deserializer[OUT]) {
	reader := kafka.NewReader(config)

	for {
		ctx, cleanup := context.WithTimeout(context.Background(), time.Minute)
		m, err := reader.ReadMessage(ctx)
		cleanup()
		if err != nil {
			panic(err)
		}

		out <- deserializer(m.Value)
	}
}

func NewKafkaDataSource[OUT interface{}](
	bootstrapUrl []string,
	topic,
	consumerId string,
	deserializer deserialize.Deserializer[OUT],
) chan OUT {
	out := make(chan OUT, 1_000)

	config := defaultConsumerConfig
	config.Brokers = bootstrapUrl
	config.Topic = topic
	config.GroupID = consumerId

	go kafkaReader[OUT](
		config,
		out,
		deserializer)

	return out
}

func KafkaDataSource[OUT interface{}](config kafka.ReaderConfig, deserializer deserialize.Deserializer[OUT]) chan OUT {
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
