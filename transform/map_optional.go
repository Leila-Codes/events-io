package transform

type MapOptionalFunction[IN, OUT interface{}] func(in IN) Optional[OUT]

func mapTransformerOpt[IN, OUT interface{}](
	input chan IN,
	transform MapOptionalFunction[IN, OUT],
	output chan OUT) {
	for {
		v := transform(<-input)
		if !v.IsEmpty() {
			output <- *v.Value
		}
	}
	close(output)
}

// MapOptional - Stateless function, similar to Map but may not always output
func MapOptional[IN, OUT interface{}](
	input chan IN,
	mapper MapOptionalFunction[IN, OUT]) chan OUT {
	out := make(chan OUT)

	go mapTransformerOpt[IN, OUT](input, mapper, out)

	return out
}
