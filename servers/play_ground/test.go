package play_ground

type DrakeStruct[T any] struct {
	Data T
}

func (d DrakeStruct[T]) getHead() T {
	return d.Data
}

func NewStruct[T any](d T) DrakeStruct[T] {
	return DrakeStruct[T]{
		Data: d,
	}
}
