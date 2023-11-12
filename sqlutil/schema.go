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

func (c Column) String() string {
	str := fmt.Sprintf("%s %s %s", c.Name, c.Type, c.Options)
	return strings.TrimSpace(str)
}

type Columns []Column

func (c Columns) Names() []string {
	names := make([]string, len(c))
	for i, column := range c {
		names[i] = column.Name
	}

	return names
}

func (c Columns) Columns() []string {
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

func (t Table) String() string {
	str := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s) %s",
		t.Name,
		strings.Join(t.Columns.Columns(), ", "),
		t.Options,
	)

	return strings.TrimSpace(str)
}

type Index struct {
	Unique  bool
	Table   string
	Columns Columns
}

func (i Index) String() string {
	indexType := "INDEX"
	if i.Unique {
		indexType = "UNIQUE INDEX"
	}

	return fmt.Sprintf(
		"CREATE %s IF NOT EXISTS %s ON %s (%s)",
		indexType,
		strings.Join(i.Columns.Names(), "_"),
		i.Table,
		strings.Join(i.Columns.Names(), ", "),
	)
}
