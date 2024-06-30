package file

import (
	"bufio"
	"log"
	"os"
)

func fileEventWriter[IN interface{}](
	input chan IN,
	serializer LineSerializer[IN],
	w *bufio.Writer,
) {
	for event := range input {
		_, err := w.Write(serializer(event))
		if err != nil {
			log.Fatal("File Sink Error - Writer Error: ", err)
		}
	}
}

func DataSink[IN interface{}](
	input chan IN,
	filePath string,
	serializer LineSerializer[IN],
) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("File Sink Error - File Error: ", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	go fileEventWriter(input, serializer, writer)
}
