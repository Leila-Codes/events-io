package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoining(t *testing.T) {
	events := make(chan string, 5)
	events <- "21"
	events <- "16"
	events <- "9"
	events <- "hello"
	events <- "world"
	close(events)

	out := Joining(events, ',')
	assert.Equal(t, out, "21,16,9,hello,world")
}
