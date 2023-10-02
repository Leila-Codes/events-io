package sink

import (
	"context"
	"github.com/Leila-Codes/events-io/sink/serialize"
	"github.com/segmentio/kafka-go"
	"time"
)

func kafkaWriter[IN interface{}](
	writer *kafka.Writer,
	in <-chan IN,
	serializer serialize.Serializer[IN]) {
	for {
		event := <-in

		ctx, cleanup := context.WithTimeout(context.Background(), 10*time.Second)
		err := writer.WriteMessages(ctx, kafka.Message{Key: nil, Value: serializer(event)})
		cleanup()
		if err != nil {
			panic(err)
		}
	}
}

const (
	defaultBackoffMin = time.Second
	defaultBackoffMax = time.Minute
)

func NewKafkaDataSink[IN interface{}](
	input <-chan IN,
	brokerURLs []string,
	topicName string,
	serializer serialize.Serializer[IN],
) {

	writer := &kafka.Writer{
		Addr:            kafka.TCP(brokerURLs...),
		Topic:           topicName,
		WriteBackoffMin: defaultBackoffMin,
		WriteBackoffMax: defaultBackoffMax,
		BatchSize:       cap(input),
		BatchBytes:      1e6, // 1 MB
		BatchTimeout:    10 * time.Second,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		RequiredAcks:    kafka.RequireNone,
		Async:           true,
		//Logger:                 nil,
		//ErrorLogger:            nil,
		AllowAutoTopicCreation: false,
	}

	kafkaWriter[IN](writer, input, serializer)
}
