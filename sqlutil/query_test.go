package sqlutil_test

import (
	"testing"

	"github.com/axatol/go-utils/sqlutil"
	"github.com/axatol/go-utils/testutil"
)

func TestValues(t *testing.T) {
	values := sqlutil.Values{"foo": "foo", "bar": 1}
	testutil.TestMany(t, []testutil.Tester{
		testutil.ElementsMatch{Expected: []any{"foo", "bar"}, Actual: values.Columns()},
		testutil.ElementsMatch{Expected: []any{":foo", ":bar"}, Actual: values.Placeholders()},
		testutil.ElementsMatch{Expected: []any{"foo = :foo", "bar = :bar"}, Actual: values.Assignment()},
	})
}

func TestInsert(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.Equal{
			Expected: "INSERT INTO table (foo, bar) VALUES (:foo, :bar)",
			Actual: sqlutil.Insert{
				Table:  "table",
				Values: sqlutil.Values{"foo": "foo", "bar": 1},
				Upsert: false,
			}.SQL(),
		},
		testutil.Equal{
			Expected: "INSERT OR REPLACE INTO table (foo, bar) VALUES (:foo, :bar) ON CONFLICT DO UPDATE SET foo = :foo, bar = :bar",
			Actual: sqlutil.Insert{
				Table:  "table",
				Values: sqlutil.Values{"foo": "foo", "bar": 1},
				Upsert: true,
			}.SQL(),
		},
	})
}

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

func TestSelect(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.Equal{
			Expected: "SELECT col0, col1, col2 FROM table",
			Actual:   sqlutil.Select{Table: "table", Columns: []string{"col0", "col1", "col2"}}.SQL(),
		},
		testutil.Equal{
			Expected: "SELECT col0, col1, col2 FROM table LIMIT 1",
			Actual:   sqlutil.Select{Table: "table", Columns: []string{"col0", "col1", "col2"}, Limit: 1}.SQL(),
		},
		testutil.Equal{
			Expected: "SELECT col0, col1, col2 FROM table WHERE ((col0 LIKE :filter0 OR col1 = :filter1) AND col2 = :filter2) LIMIT 1",
			Actual:   sqlutil.Select{Table: "table", Columns: []string{"col0", "col1", "col2"}, Limit: 1, Filter: sqlutil.And{sqlutil.Or{"col0 LIKE :filter0", "col1 = :filter1"}, "col2 = :filter2"}.SQL()}.SQL(),
		},
	})
}
