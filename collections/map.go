package collections

type Map[T any] struct {
	values map[string]T
}

func (m Map[T]) Has(key string) bool {
	_, ok := m.values[key]
	return ok
}

func (m Map[T]) Get(key string) T {
	return m.values[key]
}

func (m Map[T]) Set(key string, value T) {
	m.values[key] = value
}

func (m Map[T]) Del(key string) {
	delete(m.values, key)
}

func NewMap[T any]() Map[T] {
	values := make(map[string]T)
	return Map[T]{values}
}
