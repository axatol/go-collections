package ds

type Set[T comparable] map[T]struct{}

// Number of elements
func (s Set[T]) Len() int {
	return len(s)
}

// Calls a function for each element
func (s Set[T]) ForEach(fn func(element T) error) error {
	for v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}

	return nil
}

// Returns all entries as an array
func (s Set[T]) Entries() []T {
	index := 0
	result := make([]T, s.Len())

	s.ForEach(func(element T) error {
		result[index] = element
		index += 1
		return nil
	})

	return result
}

// Checks if the given element exists
func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}

// Add the given element to the set
func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

// Removes the given element to the set
func (s Set[T]) Del(value T) {
	delete(s, value)
}

// Creates a clone of the set and merges the elements of the given
//
// i.e. left ∪ right
func (left Set[T]) Union(right Set[T]) Set[T] {
	result := NewSet(left.Entries()...)
	result.Add(right.Entries()...)
	return result
}

// Creates a clone of the set and merges the elements of the given
//
// i.e. left ∩ right
func (left Set[T]) Intersection(right Set[T]) Set[T] {
	result := NewSet[T]()
	left.ForEach(func(element T) error {
		if right.Has(element) {
			result.Add(element)
		}

		return nil
	})

	return result
}

// Creates a clone of the set and removes the elements from the given set
//
// i.e. left - right
func (left Set[T]) Difference(right Set[T]) Set[T] {
	result := NewSet(left.Entries()...)
	right.ForEach(func(element T) error {
		result.Del(element)
		return nil
	})

	return result
}

// Instantiates a new set, optionally filling with the given elements
func NewSet[T comparable](input ...T) Set[T] {
	set := make(Set[T], len(input))
	set.Add(input...)
	return set
}
