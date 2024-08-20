package deserializer

import (
	"bytes"
	"encoding/csv"

	"github.com/Leila-Codes/events-io/transform"
)

func csvDeserializer(data []byte) []string {
	buff := bytes.NewBuffer(data)

	record, err := csv.NewReader(buff).Read()

	if err != nil {
		panic("CSV deserializer error: " + err.Error())
	}

	return record
}

func Csv(raw <-chan []byte) chan []string {
	return transform.Map(raw, csvDeserializer)
}
