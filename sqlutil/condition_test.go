package sqlutil_test

import (
	"testing"

	"github.com/axatol/go-utils/sqlutil"
	"github.com/axatol/go-utils/testutil"
)

type mockSequeliser struct{ val string }

func (s mockSequeliser) SQL() string { return s.val }

type mockStringer struct{ val string }

func (s mockStringer) String() string { return s.val }

func TestConditions(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.ElementsMatch{
			Expected: []any{"foo", "bar", "baz"},
			Actual:   sqlutil.Conditions{"foo", "bar", "baz"}.Strings(),
		},
		testutil.ElementsMatch{
			Expected: []any{"foo", "bar", "baz"},
			Actual:   sqlutil.Conditions{"foo", mockSequeliser{val: "bar"}, mockStringer{val: "baz"}}.Strings(),
		},
		testutil.Equal{
			Expected: "(foo OR bar OR baz)",
			Actual:   sqlutil.Or{"foo", "bar", "baz"}.SQL(),
		},
		testutil.Equal{
			Expected: "(foo OR (bar AND baz))",
			Actual:   sqlutil.Or{"foo", sqlutil.And{"bar", "baz"}}.SQL(),
		},
		testutil.Equal{
			Expected: "((foo AND bar) OR (baz AND qux))",
			Actual:   sqlutil.Or{sqlutil.And{"foo", "bar"}, sqlutil.And{"baz", "qux"}}.SQL(),
		},
	})
}

func TestCondition(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.Equal{
			Expected: "foo",
			Actual:   new(sqlutil.Condition).Append("foo").String(),
		},
		testutil.Equal{
			Expected: "(",
			Actual:   new(sqlutil.Condition).Open().String(),
		},
		testutil.Equal{
			Expected: ")",
			Actual:   new(sqlutil.Condition).Close().String(),
		},
		testutil.Equal{
			Expected: "AND",
			Actual:   new(sqlutil.Condition).And().String(),
		},
		testutil.Equal{
			Expected: "foo AND bar",
			Actual:   new(sqlutil.Condition).And("foo", "bar").String(),
		},
		testutil.Equal{
			Expected: "OR",
			Actual:   new(sqlutil.Condition).Or().String(),
		},
		testutil.Equal{
			Expected: "foo OR bar",
			Actual:   new(sqlutil.Condition).Or("foo", "bar").String(),
		},
		testutil.Equal{
			Expected: "col LIKE val",
			Actual:   new(sqlutil.Condition).Like("col", "val").String(),
		},
		testutil.Equal{
			Expected: "1 = 2",
			Actual:   new(sqlutil.Condition).Eq("1", "2").String(),
		},
		testutil.Equal{
			Expected: "1 > 2",
			Actual:   new(sqlutil.Condition).Gt("1", "2").String(),
		},
		testutil.Equal{
			Expected: "1 >= 2",
			Actual:   new(sqlutil.Condition).Ge("1", "2").String(),
		},
		testutil.Equal{
			Expected: "1 < 2",
			Actual:   new(sqlutil.Condition).Lt("1", "2").String(),
		},
		testutil.Equal{
			Expected: "1 <= 2",
			Actual:   new(sqlutil.Condition).Le("1", "2").String(),
		},
	})
}
