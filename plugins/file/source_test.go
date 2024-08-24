package file_test

import (
	"testing"

	"github.com/Leila-Codes/events-io/plugins/file"
	"github.com/stretchr/testify/assert"
)

func TestScannerSource(t *testing.T) {
	data := file.ScannerSource(
		3,
		"../../examples/file-example/test_data/plaintext.test.txt",
	)

	i := 0
	for msg := range data {
		switch i {
		case 0:
			assert.Equal(t, []byte("hello world"), msg)
		case 1:
			assert.Equal(t, []byte("this is a test file"), msg)
		case 2:
			assert.Equal(t, []byte("nothing magic happening here"), msg)
		}

		i++
	}

	assert.Equal(t, 3, i)
}
