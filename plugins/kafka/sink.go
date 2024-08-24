package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type MessageSerializer[IN interface{}] func(IN) kafka.Message

func DataSink[IN interface{}](
	input chan IN,
	writer *kafka.Writer,
	serializer MessageSerializer[IN],
) error {
	for msg := range input {
		err := writer.WriteMessages(
			context.TODO(),
			serializer(msg),
		)

		if err != nil {
			// log.Fatal("Kafka Sink Error - Write Error: ", err)
			return err
		}
	}

	return nil
}
