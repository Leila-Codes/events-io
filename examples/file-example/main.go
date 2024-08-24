package main

import (
	"github.com/Leila-Codes/events-io/plugins/file"
	"github.com/Leila-Codes/events-io/plugins/kafka"
	kafka2 "github.com/segmentio/kafka-go"
)

func main() {
	// read the test file line-by-line
	raw := file.ScannerSource(
		3,
		"test_data/plaintext.test.txt",
	)

	// submit each line as a new kafka event
	kafka.DataSink(
		raw,
		&kafka2.Writer{
			Addr:  kafka2.TCP("localhost:9092"),
			Topic: "test-output-1",
		},
		kafka.ToByte,
	)
}
