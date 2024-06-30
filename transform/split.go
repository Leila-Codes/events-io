package transform

import "log"

type Splitter[IN interface{}] interface {
	Splits() int
	Index(IN) int
}

type roundRobin[IN interface{}] struct {
	splits int
	iter   int
}

func (rr *roundRobin[IN]) Splits() int {
	return rr.splits
}

func (rr *roundRobin[IN]) Index(IN) int {
	curr := rr.iter

	rr.iter++

	if rr.iter >= rr.splits {
		rr.iter = 0
	}

	return curr
}

func RoundRobin[IN interface{}](splits int) Splitter[IN] {
	return &roundRobin[IN]{splits: splits}
}

func splitter[T interface{}](
	input chan T,
	outputs []chan T,
	splitter Splitter[T]) {

	for event := range input {
		idx := splitter.Index(event)

		if idx < len(outputs) {
			outputs[idx] <- event
		} else {
			log.Fatal("Split Transform Error - Index Out of Range")
		}
	}
}

func Split[T interface{}](
	input chan T,
	bufferSize uint64,
	spliterator Splitter[T],
) []chan T {
	splits := spliterator.Splits()

	outputs := make([]chan T, splits)
	for i := 0; i < splits; i++ {
		outputs[i] = make(chan T, bufferSize)
	}

	go splitter(input, outputs, spliterator)

	return outputs
}
