package sqlutil_test

import (
	"strings"
	"testing"

	"github.com/axatol/go-utils/sqlutil"
	"github.com/stretchr/testify/assert"
)

func TestColumnString(t *testing.T) {
	tests := []struct {
		input    sqlutil.Column
		expected string
	}{
		{
			input:    sqlutil.Column{Name: "name", Type: "type", Options: "options"},
			expected: "name type options",
		},
		{
			input:    sqlutil.Column{Name: "name", Type: "type"},
			expected: "name type",
		},
	}

	for _, test := range tests {
		input := test.input
		expected := test.expected
		t.Run(expected, func(t *testing.T) {
			t.Parallel()
			actual := input.SQL()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestColumnsSQLs(t *testing.T) {
	tests := []struct {
		input    sqlutil.Columns
		expected []string
	}{
		{
			input: sqlutil.Columns{
				{Name: "name", Type: "type", Options: "options"},
			},
			expected: []string{"name type options"},
		},
		{
			input: sqlutil.Columns{
				{Name: "foo", Type: "TEXT"},
				{Name: "bar", Type: "INTEGER", Options: "NOT NULL"},
			},
			expected: []string{"foo TEXT", "bar INTEGER NOT NULL"},
		},
	}

	for _, test := range tests {
		input := test.input
		expected := test.expected
		t.Run(strings.Join(expected, "_"), func(t *testing.T) {
			t.Parallel()
			actual := input.SQLs()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestTableString(t *testing.T) {
	tests := []struct {
		input    sqlutil.Table
		expected string
	}{
		{
			input: sqlutil.Table{
				Name: "table",
				Columns: sqlutil.Columns{
					{Name: "foo", Type: "TEXT"},
					{Name: "bar", Type: "INTEGER", Options: "NOT NULL"},
				},
			},
			expected: "CREATE TABLE IF NOT EXISTS table (foo TEXT, bar INTEGER NOT NULL)",
		},
		{
			input: sqlutil.Table{
				Name:    "table",
				Options: "WITHOUT ROWID",
				Columns: sqlutil.Columns{
					{Name: "foo", Type: "TEXT"},
					{Name: "bar", Type: "INTEGER", Options: "NOT NULL"},
				},
			},
			expected: "CREATE TABLE IF NOT EXISTS table (foo TEXT, bar INTEGER NOT NULL) WITHOUT ROWID",
		},
	}

	for _, test := range tests {
		input := test.input
		expected := test.expected
		t.Run(expected, func(t *testing.T) {
			t.Parallel()
			actual := input.SQL()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestIndexString(t *testing.T) {
	tests := []struct {
		input    sqlutil.Index
		expected string
	}{
		{
			input: sqlutil.Index{
				Table:   "table",
				Columns: []string{"name"},
			},
			expected: "CREATE INDEX IF NOT EXISTS name ON table (name)",
		},
		{
			input: sqlutil.Index{
				Table:   "table",
				Unique:  true,
				Columns: []string{"name", "timestamp"},
			},
			expected: "CREATE UNIQUE INDEX IF NOT EXISTS name_timestamp ON table (name, timestamp)",
		},
	}

	for _, test := range tests {
		input := test.input
		expected := test.expected
		t.Run(expected, func(t *testing.T) {
			t.Parallel()
			actual := input.SQL()
			assert.Equal(t, expected, actual)
		})
	}
}
