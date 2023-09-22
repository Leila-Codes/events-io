package transform

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	in := make(chan int)

	out := Map[int, int](in, func(in int) int {
		return in * 2
	})

	for i := 1; i <= 10; i++ {
		in <- i
		fmt.Println(<-out)
	}
}
