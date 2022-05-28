package collections

type Set[T comparable] struct {
	values map[T]bool
}

func (s Set[T]) Has(value T) bool {
	_, ok := s.values[value]
	return ok
}

func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s.values[value] = true
	}
}

func (s Set[T]) Del(value T) {
	delete(s.values, value)
}

func NewSet[T comparable](input []T) Set[T] {
	values := make(map[T]bool, len(input))
	for _, value := range input {
		values[value] = true
	}
	return Set[T]{values}
}
