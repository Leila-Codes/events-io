package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func StringValueSerializer[IN interface{}](event IN) kafka.Message {
	return kafka.Message{Value: []byte(fmt.Sprintf("%s", event))}
}

func JsonValueSerializer[IN interface{}](event IN) kafka.Message {
	data, err := json.Marshal(event)
	if err != nil {
		log.Fatal("Kafka Sink Error - Serialize Exception: ", err)
	}

	return kafka.Message{Value: data}
}
