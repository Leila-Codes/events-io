package socket

import (
	"bufio"
	"net"
)

// multiCastWriter is a struct with capabilities to "broadcast" a single event (bytes) to multiple TCP client connections.
type multiCastWriter struct {
	listeners []net.Conn
	writers   []*bufio.Writer
}

// subscribe registers a new tcp connection to the writer instance and constructs a new bufio writer for writing to the socket.
func (m *multiCastWriter) subscribe(conn net.Conn) {
	m.listeners = append(m.listeners, conn)
	m.writers = append(m.writers, bufio.NewWriter(conn))
}

// notify "broadcasts" a given slice of bytes to all currently active sockets.
func (m *multiCastWriter) notify(event []byte) {
	for _, writer := range m.writers {
		_, err := writer.Write(event)
		if err != nil {
			panic("socket multicast write failed: " + err.Error())
		}
	}
}

// ListenAndServce beings listening for new socket connections from the given listener object.
// When a connection is accepted, it is automatically subscribed to this multicast writer.
func (m *multiCastWriter) listenAndServe(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("socket server client connection failed: " + err.Error())
		}

		m.subscribe(conn)
	}
}
