package file

import (
	"bufio"
	"log"
	"os"
)

func fileEventWriter(
	input chan []byte,
	w *bufio.Writer,
) {
	for event := range input {
		_, err := w.Write(event)
		if err != nil {
			log.Fatal("File Sink Error - Writer Error: ", err)
		}
	}
}

func DataSink(
	input chan []byte,
	filePath string,
) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("File Sink Error - File Error: ", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	go fileEventWriter(input, writer)
}
