package plugins

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStream(t *testing.T) {
	data := []int{5, 19, 21, 6, 5}

	out := Stream(data)

	count := 0
	for event := range out {
		switch count {
		case 0:
			assert.Equal(t, event, 5)
		case 1:
			assert.Equal(t, event, 19)
		case 2:
			assert.Equal(t, event, 21)
		case 3:
			assert.Equal(t, event, 6)
		case 4:
			assert.Equal(t, event, 5)
		}
		count++
	}

	assert.Equal(t, count, 5)
}
