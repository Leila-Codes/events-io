package socket

import (
	"net"
)

func serverListener(
	listener net.Listener,
	delim byte,
	output chan []byte,
) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("socket server client connection failed: " + err.Error())
		}

		go socketStreamReader(conn, delim, output)
	}
}

// ServerSource is similar to the ClientSource function, only it operates as a TCP/UDP server.
// It will await for any client connections and once established a new thread is spawned to listen for
// any data packets received,
func ServerSource(
	network, host string, port int,
	delim byte,
	bufferSize uint64,
) chan []byte {
	output := make(chan []byte, bufferSize)

	listener, err := net.Listen(network, connString(host, port))
	if err != nil {
		panic("socket server source failed to start: " + err.Error())
	}

	go serverListener(listener, delim, output)

	return output
}

// ServerSink is similar to the ClientSink function, only it operates as a TCP/UDP server.
// It will await for any client connections and once established begin transmitting all events
// received from the input channel, to every client.
func ServerSink(
	input chan []byte,
	network, host string, port int,
) {
	listener, err := net.Listen(network, connString(host, port))
	if err != nil {
		panic("socket server sink failed to start: " + err.Error())
	}

	writer := &multiCastWriter{}
	go writer.listenAndServe(listener)

	for event := range input {
		writer.notify(event)
	}
}
