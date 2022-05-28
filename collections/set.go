package collections

type Set[T comparable] struct {
	Data map[T]bool
}

func (s Set[T]) Length() int {
	return len(s.Data)
}

func (s Set[T]) Entries() []T {
	index := 0
	result := make([]T, s.Length())

	for value := range s.Data {
		result[index] = value
		index += 1
	}

	return result
}

func (s Set[T]) Has(value T) bool {
	_, ok := s.Data[value]
	return ok
}

func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s.Data[value] = true
	}
}

func (s Set[T]) Del(value T) {
	delete(s.Data, value)
}

func NewSet[T comparable](input []T) Set[T] {
	values := make(map[T]bool, len(input))
	for _, value := range input {
		values[value] = true
	}
	return Set[T]{values}
}
