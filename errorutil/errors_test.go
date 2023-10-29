package errorutil_test

import (
	"errors"
	"testing"

	"github.com/axatol/go-utils/errorutil"
	"github.com/stretchr/testify/assert"
)

var (
	errFoo = errors.New("foo")
	errBar = errors.New("bar")
)

func TestJoin(t *testing.T) {
	err := errorutil.Join(errFoo, errBar)
	expected := "foo; bar"
	actual := err.Error()
	assert.Equal(t, expected, actual)
}

func TestIs(t *testing.T) {
	err := errorutil.Join(errFoo, errBar)
	assert.True(t, errorutil.Is(err, errFoo))
	assert.True(t, errorutil.Is(err, errBar))
}

func TestUnwrap(t *testing.T) {
	err := errorutil.Join(errFoo, errBar)
	expected := []error{errFoo, errBar}
	actual := errorutil.Unwrap(err)
	assert.ElementsMatch(t, expected, actual)
	// check order
	for i, expectedErr := range expected {
		actualErr := actual[i]
		assert.Equal(t, expectedErr.Error(), actualErr.Error())
	}
}
