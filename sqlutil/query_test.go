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
