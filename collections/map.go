package collections

type Map[T any] map[string]T

type MapEntry[T any] struct {
	Key   string
	Value T
}

// Number of elements
func (m Map[T]) Len() int {
	return len(m)
}

// Calls a function for each element
func (m Map[T]) ForEach(fn func(value T, key string) error) error {
	for k, v := range m {
		if err := fn(v, k); err != nil {
			return err
		}
	}

	return nil
}

// Returns the keys as an array
func (m Map[T]) Keys() []string {
	index := 0
	result := make([]string, m.Len())

	m.ForEach(func(value T, key string) error {
		result[index] = key
		index += 1
		return nil
	})

	return result
}

// Returns the values as an array
func (m Map[T]) Values() []T {
	index := 0
	result := make([]T, m.Len())

	m.ForEach(func(value T, key string) error {
		result[index] = value
		index += 1
		return nil
	})

	return result
}

// Returns an array of key-value structs of the entries
func (m Map[T]) Entries() []MapEntry[T] {
	index := 0
	result := make([]MapEntry[T], m.Len())

	m.ForEach(func(value T, key string) error {
		result[index] = MapEntry[T]{Key: key, Value: value}
		index += 1
		return nil
	})

	return result
}

// Checks if the given element exists
func (m Map[T]) Has(key string) bool {
	_, ok := m[key]
	return ok
}

// Retrieves the value associated with the given key or nil otherwise
func (m Map[T]) Get(key string) *T {
	if value, ok := m[key]; ok {
		return &value
	}

	return nil
}

// Sets the value associated with the given key, returns true if a value was overwritten
func (m Map[T]) Set(key string, value T) bool {
	exists := m.Has(key)
	m[key] = value
	return exists
}

// Deletes the value associated with the given key, returns true if a value was deleted
func (m Map[T]) Del(key string) bool {
	exists := m.Has(key)
	delete(m, key)
	return exists
}

// Instantiates a new map, optionally copying in the entries of the given map
func NewMap[T any](initial ...map[string]T) Map[T] {
	result := make(Map[T])

	if len(initial) < 1 {
		return result
	}

	for k, v := range initial[0] {
		result.Set(k, v)
	}

	return result
}
