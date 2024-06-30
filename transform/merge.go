package transform

func merger[IN interface{}](
	input chan IN,
	output chan IN,
) {
	for event := range input {
		output <- event
	}
}

func Merge[IN interface{}](
	bufferSize uint64,
	input ...chan IN,
) chan IN {
	output := make(chan IN, bufferSize)

	for _, inputChan := range input {
		go merger(inputChan, output)
	}

	return output
}
