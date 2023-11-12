package sqlutil

import (
	"fmt"
	"strings"
)

type Values map[string]any

func (v Values) Columns() []string {
	columns := make([]string, len(v))
	index := 0
	for key := range v {
		columns[index] = key
		index += 1
	}
	return columns
}

func (v Values) Placeholders() []string {
	placeholders := make([]string, len(v))
	for i, column := range v.Columns() {
		placeholders[i] = ":" + column
	}
	return placeholders
}

func (v Values) Assignment() []string {
	assignment := make([]string, len(v))
	for i, column := range v.Columns() {
		assignment[i] = fmt.Sprintf("%s = :%s", column, column)
	}
	return assignment
}

func (v Values) Values() []any {
	values := make([]any, len(v))
	index := 0
	for _, value := range v {
		values[index] = value
		index += 1
	}
	return values
}

type Insert struct {
	Table   string
	Values  Values
	Options string
	Upsert  bool
}

func (i Insert) SQL() string {
	columns := i.Values.Columns()
	placeholders := i.Values.Placeholders()
	conflict := i.Values.Assignment()

	verb := "INSERT"
	if i.Upsert {
		verb = "INSERT OR REPLACE"
	}

	query := fmt.Sprintf(
		"%s INTO %s (%s) VALUES (%s)",
		verb,
		i.Table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	if i.Options != "" {
		query = fmt.Sprintf("%s %s", query, i.Options)
	}

	if i.Upsert {
		query = fmt.Sprintf("%s ON CONFLICT DO UPDATE SET %s", query, strings.Join(conflict, ", "))
	}

	return query
}

type Conditions []any

func (c Conditions) Strings() []string {
	parts := make([]string, len(c))
	for i, expr := range c {
		switch v := expr.(type) {
		case Sequeliser:
			parts[i] = v.SQL()
		case string:
			parts[i] = v
		default:
			parts[i] = fmt.Sprint(v)
		}
	}

	return parts
}

type Or Conditions

func (o Or) SQL() string {
	exprs := Conditions(o).Strings()
	return fmt.Sprintf("(%s)", strings.Join(exprs, " OR "))
}

type And Conditions

func (a And) SQL() string {
	exprs := Conditions(a).Strings()
	return fmt.Sprintf("(%s)", strings.Join(exprs, " AND "))
}

type Select struct {
	Table   string
	Columns []string
	Filter  string
	Limit   int
}

func (s Select) SQL() string {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(s.Columns, ", "), s.Table)

	if len(s.Filter) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, s.Filter)
	}

	if s.Limit > 0 {
		query = fmt.Sprintf("%s LIMIT %d", query, s.Limit)
	}

	return query
}
