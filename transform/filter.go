package transform

type FilterFunc[IN interface{}] func(IN) bool

// Filter - Applies the given filter to each event received from input.
// If the filter returns true, the event is output to the returned channel.
// Else, the event is ignored.
func Filter[IN interface{}](input <-chan IN, filter FilterFunc[IN]) chan IN {
	out := make(chan IN)
	for {
		if event := <-input; filter(event) {
			out <- event
		}
	}
}
