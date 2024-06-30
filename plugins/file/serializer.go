package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type LineSerializer[IN interface{}] func(IN) []byte

func StringSerializer[IN interface{}](event IN) []byte {
	buff := &bytes.Buffer{}
	_, err := fmt.Fprintf(buff, "%s", event)
	if err != nil {
		log.Println("File String Serializer Error - Serializer Error: ", err)
	}
	return buff.Bytes()
}

func JsonSerializer[IN interface{}](event IN) []byte {
	data, err := json.Marshal(event)
	if err != nil {
		log.Println("File Json Serializer Error - Serializer Error: ", err)
	}

	return data
}
