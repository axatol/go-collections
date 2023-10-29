package errorutil

import "strings"

var (
	_ error       = (*errors)(nil)
	_ unwrappable = (*errors)(nil)
)

// unwrappable is the interface that wraps the Unwrap method
type unwrappable interface{ Unwrap() []error }

// errors represents zero or more errors
type errors []error

// Unwrap extracts the errors contained by Errors
func (e *errors) Unwrap() []error {
	if e != nil {
		return []error(*e)
	}

	return []error{}
}

// Error represents the collected errors as a string
func (e *errors) Error() string {
	if e == nil {
		return ""
	}

	var builder strings.Builder
	for i, err := range *e {
		if i > 0 {
			builder.WriteString("; ")
		}

		builder.WriteString(err.Error())
	}

	return builder.String()
}

// Unwrap extracts errors from an Errors struct or the error in a slice
func Unwrap(err error) []error {
	if unwrappable, ok := (err).(unwrappable); ok {
		return unwrappable.Unwrap()
	}

	return []error{err}
}

// Join collects the given list of errors into an Errors struct
func Join(err error, more ...error) error {
	errs := Unwrap(err)

	for _, err := range more {
		errs = append(errs, Unwrap(err)...)
	}

	e := errors(errs)
	return &e
}

// Is determines if err is or contains target
func Is(err, target error) bool {
	if err.Error() == target.Error() {
		return true
	}

	for _, err := range Unwrap(err) {
		if err.Error() == target.Error() {
			return true
		}
	}

	return false
}
