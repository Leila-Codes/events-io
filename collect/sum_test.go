package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	events := make(chan int, 5)
	events <- 1
	events <- 2
	events <- 3
	events <- 4
	events <- 5
	close(events)

	out := Sum(events)

	assert.Equal(t, out, 15)
}
