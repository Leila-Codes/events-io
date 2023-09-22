# Events.io
An event-streaming, functional abstraction for Golang.

## <span style="color: red;">WARNING!</span> Highly Experimental
Note that this collection of functions and tools is highly experimental. 
It's written mostly for fun to play with Generics. I have both enjoyed writing the code for this and seeing it in action when writing test applications with it (root of the repo)

## Inspiration
As a very new user to the Apache Flink framework I find myself fascinated by its abstractions and stateless + stateful functional approach.
It's a heavy inspiration for this library in Golang, using some of it's more basic principles, combined with Golang generics.

# Installation
Install the package:
```shell
go get github.com/Leila-Codes/events-io
```

# Example Application
```go
package main

import (
	"github.com/Leila-Codes/events-io/sink"
	"github.com/Leila-Codes/events-io/source"
	"github.com/segmentio/kafka-go"
	"time"
)

type MyJsonMessage struct {
	Timestamp      time.Time `json:"timestamp"`
	MessageContent string    `json:"message"`
}

func main() {
	input := source.KafkaDataSource[MyJsonMessage](
		kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "test-topic-1",
			GroupID: "my-consumer-id"},
		source.JsonDeserializer[MyJsonMessage])
	
	// TODO: Perform some transformations or processing on the data.

	sink.KafkaDataSink[MyJsonMessage](
		kafka.WriterConfig{
			Brokers: []string{"localhost:9092"},
			Topic: "my-output-topic"},
                input,
                sink.JsonSerializer[MyJsonMessage])
        }
}
```

# Further Information
## Data Input (DataSource)
DataSource is a wrapper around a continuous data stream. E.g. Kafka topic
Currently this package contains implementations for:
- File Input (**Bounded**)
- Apache Kafka Consumer (Uses kafka-go client) (**Boundless**)

## Data Output (DataSink)
DataSinks are an output for of continuous data streams. E.g. Kafka Topic
Currently this package contains implementations for:
- Apache Kafka Producer (Using kafka-go client)

**Note:** Data Sinks are currently a blocking implementation, thus they should be written last to form the main function, or otherwise you may call it as a goroutine.
```go
go sink.KafkaDataSink[MyJsonMessage](...)
```

## Transformations (Functions)
A transformer converts data from one input type to an output type. A common example is a simple "Map" function.
Due to the closure mechanism in Golang, it is possible to have both a stateless and stateful "Map" function.
See `TestTransformStream.go` to see an example of a stateful one.