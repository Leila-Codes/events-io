package transform

type KeyedEvent[IN, KEY interface{}] struct {
	Key   KEY
	Value IN
}

type KeyFunc[IN, KEY interface{}] func(IN) KEY

func KeyBy[IN, KEY interface{}](input chan IN, keyFunc KeyFunc[IN, KEY]) chan KeyedEvent[IN, KEY] {
	out := make(chan KeyedEvent[IN, KEY], 1_000)

	go keyByThread(input, keyFunc, out)

	return out
}

func keyByThread[IN, KEY interface{}](input chan IN, keyFunc KeyFunc[IN, KEY], out chan KeyedEvent[IN, KEY]) {
	for {
		event := <-input
		out <- KeyedEvent[IN, KEY]{keyFunc(event), event}
	}
}
