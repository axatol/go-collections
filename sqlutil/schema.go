package sqlutil

import (
	"fmt"
	"strings"
)

type Column struct {
	Name    string
	Type    string
	Options string
}

func (c Column) SQL() string {
	str := fmt.Sprintf("%s %s %s", c.Name, c.Type, c.Options)
	return strings.TrimSpace(str)
}

func (c Column) String() string {
	return fmt.Sprintf("%s(%s)", c.Type, c.Name)
}

type Columns []Column

func (c Columns) Names() []string {
	names := make([]string, len(c))
	for i, column := range c {
		names[i] = column.Name
	}

	return names
}

func (c Columns) SQLs() []string {
	names := make([]string, len(c))
	for i, column := range c {
		names[i] = column.SQL()
	}

	return names
}

func (c Columns) Strings() []string {
	names := make([]string, len(c))
	for i, column := range c {
		names[i] = column.String()
	}

	return names
}

type Table struct {
	Name    string
	Options string
	Columns Columns
}

func (t Table) SQL() string {
	str := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s) %s",
		t.Name,
		strings.Join(t.Columns.SQLs(), ", "),
		t.Options,
	)

	return strings.TrimSpace(str)
}

func (t Table) String() string {
	return fmt.Sprintf("TABLE(%s, [%s])", t.Name, strings.Join(t.Columns.Strings(), ", "))
}

type Index struct {
	Unique  bool
	Table   string
	Columns []string
}

func (i Index) Name() string {
	return strings.Join(i.Columns, "_")
}

func (i Index) SQL() string {
	indexType := "INDEX"
	if i.Unique {
		indexType = "UNIQUE INDEX"
	}

	return fmt.Sprintf(
		"CREATE %s IF NOT EXISTS %s ON %s (%s)",
		indexType,
		i.Name(),
		i.Table,
		strings.Join(i.Columns, ", "),
	)
}

func (i Index) String() string {
	return fmt.Sprintf("INDEX(%s, [%s])", i.Table, strings.Join(i.Columns, ", "))
}
