package file

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
)

type LineDeserializer[OUT interface{}] func([]byte) OUT

func StringDeserializer(data []byte) string {
	return string(data)
}

func JsonDeserializer[OUT interface{}](data []byte) OUT {
	out := new(OUT)
	err := json.Unmarshal(data, out)
	if err != nil {
		log.Println("JSON Source Error - Deserializer Error: ", err)
	}

	return *out
}

func CsvDeserializer(data []byte) []string {
	buff := bytes.NewBuffer(data)

	record, err := csv.NewReader(buff).Read()
	if err != nil {
		log.Fatal("CSV Source Error - Deserializer Error: ", err)
	}

	return record
}
