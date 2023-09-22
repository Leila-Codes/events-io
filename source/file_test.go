package source

import (
	"fmt"
	"github.com/Leila-Codes/events-io/source/deserialize"
	"testing"
)

func TestFileInput(t *testing.T) {
	input := FileInput[string]("test_file_input.txt", deserialize.String)

	for {
		data, ok := <-input
		if !ok {
			break
		}
		fmt.Println(data)
	}
}

type ExampleEntry struct {
	Username string `json:"username"`
	Action   string `json:"action"`
}

func TestFileInput2(t *testing.T) {
	input := FileInput[ExampleEntry]("test_json_file.json", deserialize.Json[ExampleEntry])

	for {
		data, ok := <-input
		if !ok {
			break
		}

		fmt.Printf("%+v\n", data)
	}
}
