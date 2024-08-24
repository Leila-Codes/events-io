package kafka

import (
	"context"

	"github.com/Leila-Codes/events-io/util"
	"github.com/segmentio/kafka-go"
)

func kafkaReader[OUT interface{}](
	reader *kafka.Reader,
	output chan<- OUT,
	errors chan<- error,
	transformer Reader[OUT],
) {
	for {
		m, err := reader.ReadMessage(context.TODO())

		if err != nil {
			util.MustWriteError(err, errors)
		}

		output <- transformer(m)
	}
}

func DataSource[OUT interface{}](
	config kafka.ReaderConfig,
	bufferSize uint64,
	transformer Reader[OUT],
) (chan OUT, chan error) {

	var (
		out    = make(chan OUT, bufferSize)
		err    = make(chan error)
		reader = kafka.NewReader(config)
	)

	go kafkaReader(reader, out, err, transformer)

	return out, err
}
