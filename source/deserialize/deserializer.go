package deserialize

type Deserializer[OUT interface{}] func([]byte) OUT

func String(in []byte) string {
	return string(in)
}
