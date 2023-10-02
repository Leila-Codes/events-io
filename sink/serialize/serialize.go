package serialize

type Serializer[IN interface{}] func(IN) []byte

type Stringable interface {
	String() string
}

func String[IN Stringable](in IN) []byte {
	return []byte(in.String())
}
