package serializer

import (
	"bytes"
	"encoding/csv"

	"github.com/Leila-Codes/events-io/transform"
)

func csvSerializer(input []string) []byte {
	buff := &bytes.Buffer{}
	writer := csv.NewWriter(buff)

	err := writer.Write(input)
	if err != nil {
		panic("CSV serializer error: " + err.Error())
	}
	writer.Flush()

	return buff.Bytes()
}

func Csv(input <-chan []string) chan []byte {
	return transform.Map(input, csvSerializer)
}
