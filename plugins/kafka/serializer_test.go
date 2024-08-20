package kafka_test

import (
	"testing"

	"github.com/Leila-Codes/events-io/plugins/kafka"
	kafka2 "github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	source := make(chan string, 3)

	source <- "hello world"
	source <- "how are you?"
	source <- "goodbye"

	close(source)

	i := 0
	for event := range source {
		msg := kafka.ToString(event)

		assert.IsType(t, kafka2.Message{}, msg, "should produce kafka message")
		switch i {
		case 0:
			assert.Equal(t, []byte("hello world"), msg.Value)
		case 1:
			assert.Equal(t, []byte("how are you?"), msg.Value)
		case 2:
			assert.Equal(t, []byte("goodbye"), msg.Value)
		}

		i++
	}

	assert.Equal(t, i, 3)
}

func TestToByte(t *testing.T) {
	source := make(chan []byte, 2)

	source <- []byte("hello world")
	source <- []byte("goodbye")

	close(source)

	i := 0
	for event := range source {
		msg := kafka.ToByte(event)

		assert.IsType(t, kafka2.Message{}, msg, "should produce kafka message")
		switch i {
		case 0:
			assert.Equal(t, []byte("hello world"), msg.Value)
		case 1:
			assert.Equal(t, []byte("goodbye"), msg.Value)
		}

		i++
	}

	assert.Equal(t, i, 2)
}

type ComplexEvent struct {
	Key, Value []byte
}

func TestToKeyValue(t *testing.T) {
	getKey := func(e ComplexEvent) []byte {
		return e.Key
	}

	getValue := func(e ComplexEvent) []byte {
		return e.Value
	}

	source := make(chan ComplexEvent, 2)
	source <- ComplexEvent{Key: []byte("hello"), Value: []byte("world")}
	source <- ComplexEvent{Key: []byte("good"), Value: []byte("bye")}
	close(source)

	i := 0
	for event := range source {
		msg := kafka.ToKeyValue(event, getKey, getValue)

		assert.IsType(t, kafka2.Message{}, msg)

		switch i {
		case 0:
			assert.Equal(t, msg.Key, []byte("hello"))
			assert.Equal(t, msg.Value, []byte("world"))
		case 1:
			assert.Equal(t, msg.Key, []byte("good"))
			assert.Equal(t, msg.Value, []byte("bye"))
		}

		i++
	}

	assert.Equal(t, i, 2)
}
