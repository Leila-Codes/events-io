package socket

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func connString(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func socketStreamReader(
	socket net.Conn,
	delim byte,
	output chan []byte,
) {

	reader := bufio.NewReader(socket)

	for {
		data, err := reader.ReadSlice(delim)

		output <- data

		if err == io.EOF {
			close(output)
			break
		} else if err != nil {
			panic("tcp reader error: " + err.Error())
		}
	}
}

func socketStreamWriter(
	input chan []byte,
	conn net.Conn,
) {
	writer := bufio.NewWriter(conn)

	for event := range input {
		_, err := writer.Write(event)
		if err != nil {
			panic("socket stream writer error: " + err.Error())
		}
	}
}

// SocketSource represents a TCP/UDP client connection.
// Once established the socket is continuously polled and read until
// the delim byte is encountered, queuing all bytes read up to this
// point onto the output channel.
func ClientSource(
	network string,
	host string, port int,
	delim byte,
	bufferSize uint64,
) chan []byte {

	output := make(chan []byte, bufferSize)

	conn, err := net.Dial(network, connString(host, port))
	if err != nil {
		panic("socket source connect failed: " + err.Error())
	}

	go socketStreamReader(conn, delim, output)

	return output
}

// DataSink represents a TCP/UDP client connection. Once established,
// all byte events received from the input channel are transmitted over the socket.
func ClientSink(
	input chan []byte,
	network string,
	host string, port int,
) {
	conn, err := net.Dial(network, connString(host, port))
	if err != nil {
		panic("socket sink connect failed: " + err.Error())
	}

	socketStreamWriter(input, conn)
}
