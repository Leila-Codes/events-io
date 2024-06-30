package transform

type MapFunction[IN, OUT interface{}] func(in IN) OUT

func mapTransformer[IN, OUT interface{}](
	input chan IN,
	transform MapFunction[IN, OUT],
	output chan OUT,
) {
	for {
		output <- transform(<-input)
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
