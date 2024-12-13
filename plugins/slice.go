package plugins

func sliceStream[IN interface{}](input []IN, out chan<- IN) {
	for _, item := range input {
		out <- item
	}
	close(out)
}

func Stream[IN interface{}](input []IN) <-chan IN {
	out := make(chan IN)

	go sliceStream(input, out)

	return out
}
