# Events.io
An event-streaming, functional abstraction for Golang.

## Inspiration
Inspiration came from the Apache Flink framework and it's stateful+stateless functions approach to event-handling. Other elements such as AWS Lambda as well. This library attempts to provide a variety of simple functions to skeleton out your streaming application or microservice, supported heavily by Golang generics.

# Installation
1. Install the main package:
```shell
go get github.com/Leila-Codes/events-io
```
2. Install any desired plugins, e.g. for kafka
```shell
go get github.com/Leila-Codes/events-io/plugins/kafka
```

# Examples
Examples can be found under the `./examples` directory in this repository.

# API Principles
## Data Input (DataSource)
A Data Source is a functio that returns a stream of input data. Streams can be boundless (i.e. kafka or SQL) or it can be bounded (e.g. a file)
Below lists the currently maintained plugins and their function definition for you to use:
- File Input 

  💾 `github.com/Leila-Codes/events-io/plugins/file`
    ```go
    file.ScannerSource(bufferSize uint64, filePaths... string)
    ```
- Apache Kafka Consumer (Wraps around `github.com/segmentio/kafka-go`) 
  
  💾 `github.com/Leila-Codes/events-io/plugins/kafka`
    ```go
    kafka.DataSource(
      kafka2.ReaderConfig{}, 
      bufferSize uint64, 
      kafka.Reader, // e.g. kafka.RawMessages | kafka.ByteValue | kafka.StringValue
    )
    ```
- Socket (Client and server supported) connection (e.g. TCP or UDP)

  💾 `github.com/Leila-Codes/events-io/plugins/socket`

  - As Client
    ```go
    socket.ClientSource(
      network string, // "tcp" or "udp"
      host string, port int,
      delim byte, // char to read until before returning as event on channel.
      bufferSize uint64,
    )
    ```

  - As Server
    ```go
    socket.ServerSource(
      network string, // "tcp" or "udp"
      host string, port int,
      delim byte, // char to read until before returning as event on channel.
      bufferSize uint64,
    )
    ```

- SQL Table (supports PostgreSQL, MS-SQL or MySQL)
  
  💾 `github.com/Leila-Codes/events-io/plugins/sql_io`
    ```go
    sql_io.DataSource(
        input <-chan IN,
        driverName, connString string,
        selectStmt string,
        setter SqlParamSetter[IN], // function recieving IN type and returning params.
        scanner SqlScanner[OUT], // function that receives sql.Row and return a new OUT type.
        bufferSize uint64,
    )
    ```

## Data Output (DataSink)
DataSinks are a synchronous (blocking) output, generally the destination of your data stream at the end.
Below lists the currently maintained plugins and their function definition for you to use:
- File Output 

  💾 `github.com/Leila-Codes/events-io/plugins/file`
  ```go
    file.DataSink(
      input chan []byte, // expects bytes, serialize with `transform/serializer` package.
      filePath string,
    )`

  ```
- Apache Kafka Producer (Wraps around `github.com/segmentio/kafka-go` client) 

  💾 `github.com/Leila-Codes/events-io/plugins/kafka`
  ```go
    kafka.DataSink(
      input chan, // channel of events, usually either []byte or string
      *kafka.Writer, // configured segmentio/kafka-go
      kafka.Builder, // typically kafka.ToByte | kafka.ToKeyValue | kafka.ToString
    )
  ```

- Socket (Client and server supported) (E.g. TCP or UDP)

  💾 `github.com/Leila-Codes/events-io/plugins/socket`
  
  - As Client
    ```go
    socket.ClientSink(
      input chan []byte, // expects bytes, use transform.serializer package.
      network string, // "tcp" or "udp"
      host string, port int,
    )
    ```

  - As Server \
    *writes every single input message to every single client that connects.*
    ```go
    socket.ServerSink(
      input chan []byte, // expects bytes, use transform.serializer package
      network string, // "tcp" or "udp"
      host string, port int,
    )
    ```

- SQL Table (MySQL, PostgrSQL and MS-SQL supported)

  💾 `github.com/Leila-Codes/events-io/plugins/sql_io`
  ```go
  sql_io.DataSink(
      input chan, 
      driverName, connString,
      insertStmt string, // e.g. INSERT INTO MyTable (ID, Username, Email) VALUES (?1, ?2, ?3)
      valuer sql_io.SqlValuer[IN], // function to convert IN type into list of sql.Driver values ([]interface{})
      batchSize int, // max # rows in a single insert
      batchTimeout time.Duration, // max time to wait for batchSize or else flush anyway
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
 - **KeyBy** \
  `KeyBy[IN, KEY](input chan IN, KeyFunc[IN] KEY) chan KeyedEvent[IN, KEY]`

# Serialization
Most of the input/output plugins generally expect data to be received or sent as `[]byte`. A subsidiary to the `transform` package, provides `serializer` and `deserializer` packages. Implementations provided are listed below:
- **Json** \
  Json support is provided in both packages, to deserialize []byte to a struct of your choosing and re-serializer to []byte later.
  ```go
  // events chan OUT
  events := deserializer.Json[OUT interface{}](input chan []byte)

  // output chan []byte
  output := serializer.Json(input chan any)
  ```
- **Csv** \
  Csv support is provided in both packages, to deserialize []byte to a struct of your choosing and re-serializer to []byte later.
  ```go
  // rows chan []string
  rows := deserializer.Csv(input chan []byte) 

  // output chan []byte
  output := serializer.Csv(input chan []string)
  ```
  
- **Gob** \
  Gob support is provided in both packages, to deserialize []byte to a struct of your choosing and re-serializer to []byte later.
  ```go
  // events chan []OUT
  events := deserializer.Gob[OUT interface{}](input chan []byte) 

  // output chan []byte
  output := serializer.Gob(input chan any)
  ```