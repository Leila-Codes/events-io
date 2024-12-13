package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	events := make(chan int, 5)
	events <- 5
	events <- 19
	events <- 15
	events <- 12
	events <- 21
	close(events)

	out := Slice(events)

	assert.Len(t, out, 5)
	assert.Equal(t, out, []int{5, 19, 15, 12, 21})
}
