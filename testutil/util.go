package testutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Tester interface {
	Test(t *testing.T)
}

func TestMany(t *testing.T, tests []Tester) {
	for i, test := range tests {
		test := test
		t.Run(fmt.Sprint(i), test.Test)
	}
}

type Equal struct{ Expected, Actual any }

func (test Equal) Test(t *testing.T) {
	t.Parallel()
	assert.Equal(t, test.Expected, test.Actual)
}

type ElementsMatch struct{ Expected, Actual any }

func (test ElementsMatch) Test(t *testing.T) {
	t.Parallel()
	assert.ElementsMatch(t, test.Expected, test.Actual)
}
