package transform

type MapFunction[IN, OUT interface{}] func(in IN) OUT
type MapOptionalFunction[IN, OUT interface{}] func(in IN) Optional[OUT]

func mapTransformer[IN, OUT interface{}](
	input chan IN,
	transform MapFunction[IN, OUT],
	output chan OUT,
) {
	for {
		output <- transform(<-input)
	}
}

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
}

// Map - Stateless function, calls mapper to transform data from IN type to OUT type.
func Map[IN, OUT interface{}](
	input chan IN,
	mapper MapFunction[IN, OUT]) chan OUT {

	out := make(chan OUT)

	go mapTransformer[IN, OUT](input, mapper, out)

	return out
}

// MapOptional - Stateless function, similar to Map but may not always output
func MapOptional[IN, OUT interface{}](
	input chan IN,
	mapper MapOptionalFunction[IN, OUT]) chan OUT {
	out := make(chan OUT)

	go mapTransformerOpt[IN, OUT](input, mapper, out)

	return out
}
