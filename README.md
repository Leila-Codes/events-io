# Events.io
An event-streaming, functional abstraction for Golang.

## Inspiration
As a very new user to the Apache Flink framework I find myself fascinated by its abstractions and stateless + stateful functional approach.
It's a heavy inspiration for this library in Golang, using some of it's more basic principles, combined with Golang generics.

# Installation
Install the package:
```shell
go get github.com/Leila-Codes/events-io
```

# Examples
Examples can be found under the `./examples` directory in this repository. Additionally you can pull these with go:
    `go get github.com/Leila-Codes/events-io/examples`

# Further Information
## Data Input (DataSource)
DataSource is a wrapper around a continuous (or bounded) data stream. E.g. Kafka topic, SQL table or one or more files.
Currently this package and it's plugins provides functions for:
- File Input
    ```go
    file.DataSource(file.LineDeserializer, bufferSize uint64, filePaths... string)
    ```
- Apache Kafka Consumer (Using segmentio/kafka-go client)
    ```go
    kafka.DataSource(kafka.ReaderConfig, bufferSize uint64, kafka.KafkaDeserializer)
    ```
- SQL Table (PostgreSQL only, using lib/pq client)
    ```go
    sql_io.DataSource(
        input <-chan IN,
        driverName, connString string,
        selectStmt string,
        setter SqlParamSetter[IN],
        scanner SqlScanner[OUT],
        bufferSize uint64,
    )
    ```

## Data Output (DataSink)
DataSinks are an output for of continuous (or bounded) data streams.
Currently this package and it's plugins provides functions for:
- Apache Kafka Producer (Using segmentio/kafka-go client) \
  ```
    kafka.DataSink(input chan, *kafka.Writer, kafka.MessageSerializer)
  ```
- File Output 
  ```
    file.DataSink(input chan, filePath string)`
  ```
- SQL Table (PostgreSQL only, using lib/pq client) \
  ```
  sql_io.DataSink(
      input chan, 
      driverName, connString, insertStmt string,
      valuer sql_io.SqlValuer[IN],
      batchSize int,
      batchTimeout time.Duration,
  )
  ```

## Transformations (Functions)
Transformers allow manipulating your events as they are received from data sources.
These manipulations may include; filtering out certain events, grouping related events together and mapping the data from one structure to another.

**Provided functions are**
 - **Filter** \
   `Filter[T](input chan T, FilterFunc[T])` \
   _Where FilterFunc has the signature `func(T) bool`_ \
    `true` = output the message, `false` = ignore it. 
 - **Map** \
  `Map[IN, OUT](input chan IN, MapFunc[IN, OUT])` \
   Where MapFunc has the signature `func(IN) OUT` the new / returned value is emitted on the output channel.
 - **MapOptional** \
  `MapOptional[IN, OUT](input chan IN, MapFuncOptional[IN, OUT])` \
    Where MapFuncOptional has the signature `func(IN) Optional[OUT]`, same as above, but Optional.Empty is ignored (not output).
 - **Split** \
  `Split[T](input chan T, Splitter[T]) []chan T` \
    Where Splitter implements the interface `Splitter[T]` each event is passed to a respective output channel. `RoundRobin(N)` can be used for simple round-robin split. 
 - **Merge** \
  `Merge[T](input... chan T) chan T` \
    Merging events from all channels into the output.

    *Still Under development:*
 - **KeyBy** \
  `KeyBy[IN, KEY](input chan IN, KeyFunc[IN] KEY) chan KeyedEvent[IN, KEY]`