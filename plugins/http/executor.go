package http

import "net/http"

func requestExecutor[T, OUT interface{}](
	input <-chan T,
	builder RequestBuilder[T],
	parser ResponseParser[OUT],
	errors chan<- error,
	output chan OUT,
) {
	for event := range input {
		req, err := builder(event)
		if err != nil {
			errors <- err
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			errors <- err
			return
		}

		output <- parser(res)
	}
}
