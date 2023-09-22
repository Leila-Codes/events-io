package source

import (
	"bufio"
	"github.com/Leila-Codes/events-io/source/deserialize"
	"os"
)

func fileReader[OUT interface{}](filePath string, deserializer deserialize.Deserializer[[]byte, OUT], output chan OUT) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		output <- deserializer(reader.Bytes())
	}

	close(output)
	f.Close()
}

func FileInput[OUT interface{}](
	filePath string,
	deserializer deserialize.Deserializer[[]byte, OUT]) chan OUT {

	output := make(chan OUT)

	go fileReader[OUT](filePath, deserializer, output)

	return output
}
