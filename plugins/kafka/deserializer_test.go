package kafka_test

import (
	"testing"
	"time"

	"github.com/Leila-Codes/events-io/plugins/kafka"
	kafka2 "github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

var (
	desrializeTestMsg1  = kafka2.Message{Key: []byte("hello"), Value: []byte("world")}
	deserializeTestMsg2 = kafka2.Message{Key: []byte("good"), Value: []byte("bye")}
)

func TestRawMessage(t *testing.T) {
	var (
		out1 = kafka.RawMessage(desrializeTestMsg1)
		out2 = kafka.RawMessage(deserializeTestMsg2)
	)

	assert.IsType(t, kafka2.Message{}, out1)
	assert.IsType(t, kafka2.Message{}, out2)

	assert.Equal(t, []byte("hello"), out1.Key)
	assert.Equal(t, []byte("world"), out1.Value)

	assert.Equal(t, []byte("good"), out2.Key)
	assert.Equal(t, []byte("bye"), out2.Value)
}

func TestByteValue(t *testing.T) {

	var (
		out1 = kafka.ByteValue(desrializeTestMsg1)
		out2 = kafka.ByteValue(deserializeTestMsg2)
	)

	assert.IsType(t, []byte{}, out1, "should produce message value as BYTES")
	assert.IsType(t, []byte{}, out2, "should produce message value as BYTES")

	assert.Equal(t, []byte("world"), out1, "should contain the kafka message value only")
	assert.Equal(t, []byte("bye"), out2, "should contain the kafka message value only")
}

func TestStringValue(t *testing.T) {

	var (
		out1 = kafka.StringValue(desrializeTestMsg1)
		out2 = kafka.StringValue(deserializeTestMsg2)
	)

	assert.IsType(t, "", out1, "should produce message value as STRING")
	assert.IsType(t, "", out2, "should produce message value as STRING")

	assert.Equal(t, "world", out1, "should contain the correct kafka message value")
	assert.Equal(t, "bye", out2, "should contain the correct kafka message value")
}

func TestTimestampsOnly(t *testing.T) {
	deserializeTestMsg3 := kafka2.Message{Time: time.Date(2024, time.August, 20, 20, 6, 0, 0, time.UTC)}

	out1 := kafka.TimestampsOnly(deserializeTestMsg3)

	assert.IsType(t, time.Time{}, out1, "should produce message timestamp as time.Time")

	assert.Equal(t, time.Date(2024, time.August, 20, 20, 6, 0, 0, time.UTC), out1)
}
