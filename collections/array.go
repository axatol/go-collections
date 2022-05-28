package collections

type Array[T any] struct {
	values []T
}

func (a Array[T]) ForEach(fn func(T, int)) {
	for index, element := range a.values {
		fn(element, index)
	}
}

func (a Array[T]) Map(fn func(T, int) interface{}) []interface{} {
	result := make([]interface{}, len(a.values))

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

func NewArray[T any](values []T) Array[T] {
	return Array[T]{values}
}
