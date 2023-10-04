package ds

type Map[K comparable, V any] map[K]V

type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// Number of elements
func (m Map[K, V]) Len() int {
	return len(m)
}

// Calls a function for each element
func (m Map[K, V]) ForEach(fn func(value V, key K) error) error {
	for k, v := range m {
		if err := fn(v, k); err != nil {
			return err
		}
	}

	return nil
}

// Returns the keys as an array
func (m Map[K, V]) Keys() []K {
	index := 0
	result := make([]K, m.Len())

	m.ForEach(func(value V, key K) error {
		result[index] = key
		index += 1
		return nil
	})

	return result
}

// Returns the values as an array
func (m Map[K, V]) Values() []V {
	index := 0
	result := make([]V, m.Len())

	m.ForEach(func(value V, key K) error {
		result[index] = value
		index += 1
		return nil
	})

	return result
}

// Returns an array of key-value structs of the entries
func (m Map[K, V]) Entries() []MapEntry[K, V] {
	index := 0
	result := make([]MapEntry[K, V], m.Len())

	m.ForEach(func(value V, key K) error {
		result[index] = MapEntry[K, V]{Key: key, Value: value}
		index += 1
		return nil
	})

	return result
}

// Checks if the given element exists
func (m Map[K, V]) Has(key K) bool {
	_, ok := m[key]
	return ok
}

// Retrieves the value associated with the given key or nil otherwise
func (m Map[K, V]) Get(key K) *V {
	if value, ok := m[key]; ok {
		return &value
	}

	return nil
}

// Sets the value associated with the given key, returns true if a value was overwritten
func (m Map[K, V]) Set(key K, value V) bool {
	exists := m.Has(key)
	m[key] = value
	return exists
}

// Deletes the value associated with the given key, returns true if a value was deleted
func (m Map[K, V]) Del(key K) bool {
	exists := m.Has(key)
	delete(m, key)
	return exists
}

// Instantiates a new map, optionally copying in the entries of the given map
func NewMap[K comparable, V any](initial ...map[K]V) Map[K, V] {
	result := make(Map[K, V])

	if len(initial) < 1 || initial[0] == nil {
		return result
	}

	for k, v := range initial[0] {
		result.Set(k, v)
	}

	return result
}
