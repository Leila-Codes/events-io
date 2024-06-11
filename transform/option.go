package transform

type Optional[T interface{}] struct {
	Value *T
}

func (op Optional[T]) IsEmpty() bool {
	return op.Value == nil
}

func Empty[T interface{}]() Optional[T] {
    return Optional[T]{
        Value: nil,
    }
}