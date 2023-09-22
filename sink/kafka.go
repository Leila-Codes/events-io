package sink

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func kafkaWriter[IN interface{}](
	writer *kafka.Writer,
	in chan IN,
	serializer ValueSerializer[IN]) {
	for {
		event := <-in

		ctx, cleanup := context.WithTimeout(context.Background(), 10*time.Second)
		err := writer.WriteMessages(ctx, serializer(event))
		cleanup()
		if err != nil {
			panic(err)
		}
	}
}

func KafkaDataSink[IN interface{}](
	config kafka.WriterConfig,
	in chan IN,
	serializer ValueSerializer[IN]) {
	writer := &kafka.Writer{
		Addr:        kafka.TCP(config.Brokers...),
		Topic:       config.Topic,
		Balancer:    config.Balancer,
		MaxAttempts: config.MaxAttempts,
		//WriteBackoffMin:        0,
		//WriteBackoffMax:        0,
		BatchSize:              config.BatchSize,
		BatchBytes:             int64(config.BatchBytes),
		BatchTimeout:           config.BatchTimeout,
		ReadTimeout:            config.ReadTimeout,
		WriteTimeout:           config.WriteTimeout,
		RequiredAcks:           kafka.RequiredAcks(config.RequiredAcks),
		Async:                  true,
		AllowAutoTopicCreation: false,
	}

	kafkaWriter[IN](writer, in, serializer)
}
