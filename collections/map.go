package collections

type Map[T any] struct {
	Data map[string]T
}

type MapEntry[T any] struct {
	Key   string
	Value T
}

func (m Map[T]) Length() int {
	return len(m.Data)
}

func (m Map[T]) Keys() []string {
	index := 0
	result := make([]string, m.Length())

	for value := range m.Data {
		result[index] = value
		index += 1
	}

	return result
}

func (m Map[T]) Values() []T {
	index := 0
	result := make([]T, m.Length())

	for _, value := range m.Data {
		result[index] = value
		index += 1
	}

	return result
}

func (m Map[T]) Entries() []MapEntry[T] {
	index := 0
	result := make([]MapEntry[T], m.Length())

	for key, value := range m.Data {
		result[index] = MapEntry[T]{Key: key, Value: value}
		index += 1
	}

	return result
}

func (m Map[T]) Has(key string) bool {
	_, ok := m.Data[key]
	return ok
}

func (m Map[T]) Get(key string) T {
	return m.Data[key]
}

func (m Map[T]) Set(key string, value T) {
	m.Data[key] = value
}

func (m Map[T]) Del(key string) {
	delete(m.Data, key)
}

func NewMap[T any]() Map[T] {
	values := make(map[string]T)
	return Map[T]{values}
}
