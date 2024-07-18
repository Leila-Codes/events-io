package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataSource_Plaintext(t *testing.T) {
	data := DataSource(
		StringDeserializer,
		1000,
		"tests/plaintext.test.txt",
	)

	i := 0
	for msg := range data {
		i++

		switch i {
		case 1:
			assert.Equal(t, "hello world", msg)
		case 2:
			assert.Equal(t, "this is a test file", msg)
		case 3:
			assert.Equal(t, "nothing magic happening here", msg)
		}
	}
	assert.Equal(t, 3, i)
}

func TestDataSource_CSV(t *testing.T) {
	data := DataSource(
		CsvDeserializer,
		1000,
		"tests/csv.test.csv",
	)

	i := 0
	for msg := range data {
		i++

		switch i {
		case 1:
			assert.Equal(t, []string{"username", "message", "timestamp"}, msg)
		case 2:
			assert.Equal(t, []string{"leila-codes", "hello world", "2024-07-28T20:52:00"}, msg)
		case 3:
			assert.Equal(t, []string{"fallflowers", "testy test", "2024-07-28T20:53:12"}, msg)
		}
	}

	assert.Equal(t, 3, i)
}
