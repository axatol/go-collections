package collections

type Array[T any] struct {
	Data []T
}

func (a Array[T]) Length() int {
	return len(a.Data)
}

func (a Array[T]) ForEach(fn func(T, int)) {
	for index, element := range a.Data {
		fn(element, index)
	}
}

func (a Array[T]) Map(fn func(T, int) interface{}) []interface{} {
	result := make([]interface{}, len(a.Data))

	a.ForEach(func(element T, index int) {
		result[index] = fn(element, index)
	})

	return result
}

func (a Array[T]) Filter(fn func(T, int) bool) Array[T] {
	result := []T{}

	a.ForEach(func(element T, index int) {
		if fn(element, index) {
			result = append(result, element)
		}
	})

	return Array[T]{result}
}

func (a Array[T]) Find(fn func(T, int) bool) (*T, int) {
	for i, el := range a.Data {
		if fn(el, i) {
			return &el, i
		}
	}

	return nil, 0
}

func NewArray[T any](values []T) Array[T] {
	return Array[T]{values}
}
