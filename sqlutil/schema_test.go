package sqlutil_test

import (
	"testing"

	"github.com/axatol/go-utils/sqlutil"
	"github.com/axatol/go-utils/testutil"
)

func TestColumnString(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.Equal{
			Expected: "name type options",
			Actual:   sqlutil.Column{Name: "name", Type: "type", Options: "options"}.SQL(),
		},
		testutil.Equal{
			Expected: "name type",
			Actual:   sqlutil.Column{Name: "name", Type: "type"}.SQL(),
		},
	})
}

func TestColumnsSQLs(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.Equal{
			Expected: []string{"name type options"},
			Actual: sqlutil.Columns{
				{Name: "name", Type: "type", Options: "options"},
			}.SQLs(),
		},
		testutil.Equal{
			Expected: []string{"foo TEXT", "bar INTEGER NOT NULL"},
			Actual: sqlutil.Columns{
				{Name: "foo", Type: "TEXT"},
				{Name: "bar", Type: "INTEGER", Options: "NOT NULL"},
			}.SQLs(),
		},
	})
}

func TestTableString(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.Equal{
			Expected: "CREATE TABLE IF NOT EXISTS table (foo TEXT, bar INTEGER NOT NULL)",
			Actual: sqlutil.Table{
				Name: "table",
				Columns: sqlutil.Columns{
					{Name: "foo", Type: "TEXT"},
					{Name: "bar", Type: "INTEGER", Options: "NOT NULL"},
				},
			}.SQL(),
		},
		testutil.Equal{
			Expected: "CREATE TABLE IF NOT EXISTS table (foo TEXT, bar INTEGER NOT NULL) WITHOUT ROWID",
			Actual: sqlutil.Table{
				Name:    "table",
				Options: "WITHOUT ROWID",
				Columns: sqlutil.Columns{
					{Name: "foo", Type: "TEXT"},
					{Name: "bar", Type: "INTEGER", Options: "NOT NULL"},
				},
			}.SQL(),
		},
	})
}

func TestIndexString(t *testing.T) {
	testutil.TestMany(t, []testutil.Tester{
		testutil.Equal{
			Expected: "CREATE INDEX IF NOT EXISTS name ON table (name)",
			Actual: sqlutil.Index{
				Table:   "table",
				Columns: []string{"name"},
			}.SQL(),
		},
		testutil.Equal{
			Expected: "CREATE UNIQUE INDEX IF NOT EXISTS name_timestamp ON table (name, timestamp)",
			Actual: sqlutil.Index{
				Table:   "table",
				Unique:  true,
				Columns: []string{"name", "timestamp"},
			}.SQL(),
		},
	})
}
