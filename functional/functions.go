package fp

type ForEachFn[T any] func(element T, index int) error

// Executes fn for each element in the array
//
// Returns early if an error is returned
func ForEach[T any](inputs []T, fn ForEachFn[T]) error {
	var err error
	for index, element := range inputs {
		if err = fn(element, index); err != nil {
			return err
		}
	}

	return nil
}

type MapFn[Input, Output any] func(element Input, index int) (Output, error)

// Executes mapper for each element, returning the result
//
// Returns early if an error is returned
func Map[Input, Output any](inputs []Input, mapper MapFn[Input, Output]) ([]Output, error) {
	result := make([]Output, len(inputs))

	err := ForEach(inputs, func(element Input, index int) error {
		var err error
		if result[index], err = mapper(element, index); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

type ReduceFn[Accumulator, T any] func(accumulator Accumulator, element T, index int) (Accumulator, error)

// Executes reducer for each element, the result being included in the parameter of the next execution
//
// Returns early if an error is given
func Reduce[Accumulator, T any](inputs []T, initializer Accumulator, reducer ReduceFn[Accumulator, T]) (*Accumulator, error) {
	accumulator := initializer

	err := ForEach(inputs, func(element T, index int) error {
		var err error
		if accumulator, err = reducer(accumulator, element, index); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &accumulator, nil
}

type FilterFn[T any] func(element T, index int) bool

// Returns an array of elements for which the filter evaluated to true
func Filter[T any](inputs []T, filter FilterFn[T]) []T {
	result := []T{}

	ForEach(inputs, func(input T, index int) error {
		if filter(input, index) {
			result = append(result, input)
		}

		return nil
	})

	return result
}

// Returns the first element for which the filter evaluated to true or nil if none found
func Find[T any](inputs []T, filter FilterFn[T]) *T {
	for i, v := range inputs {
		if filter(v, i) {
			return &v
		}
	}

	return nil
}

// Flattens nested arrays
func Flat[T any](inputs [][]T) []T {
	result := []T{}

	for _, v := range inputs {
		result = append(result, v...)
	}

	return result
}
