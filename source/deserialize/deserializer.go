package deserialize

type Deserializer[IN, OUT interface{}] func(IN) OUT

func String(in []byte) string {
	return string(in)
}
