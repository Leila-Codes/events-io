package collect

import "cmp"

func Sum[IN cmp.Ordered](input chan IN) IN {
	var val IN

	for event := range input {
		val += event
	}

	return val
}
