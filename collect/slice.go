package collect

func Slice[IN interface{}](input chan IN) []IN {
	out := make([]IN, 0)

	for event := range input {
		out = append(out, event)
	}

	return out
}
