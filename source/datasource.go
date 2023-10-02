package source

import "github.com/Leila-Codes/events-io/source/deserialize"

type DataSource[IN, CONF, OUT interface{}] func(config CONF, deserializer deserialize.Deserializer[OUT]) chan OUT

// BoundedDataSource - Represents a stream of bounded data (data with an end)
// Returns channel of "events" read from the source, and a "done" channel.
type BoundedDataSource[IN, CONF, OUT interface{}] func(config CONF, deserializer deserialize.Deserializer[OUT]) (chan OUT, chan struct{})
