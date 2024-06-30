package file

import (
	"bufio"
	"log"
	"os"
)

func fileReader[OUT interface{}](
	deserializer LineDeserializer[OUT],
	output chan OUT,
	filePath string,
) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("File Data Source - File Error: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		output <- deserializer(scanner.Bytes())
	}
}

func multiFileReader[OUT interface{}](
	deserializer LineDeserializer[OUT],
	output chan OUT,
	filePaths ...string,
) {
	for _, filePath := range filePaths {
		fileReader[OUT](deserializer, output, filePath)
	}
}

func DataSource[OUT interface{}](
	deserializer LineDeserializer[OUT],
	bufferSize uint64,
	filePaths ...string,
) chan OUT {
	out := make(chan OUT, bufferSize)

	go multiFileReader(deserializer, out, filePaths...)

	return out
}
