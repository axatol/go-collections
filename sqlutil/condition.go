package sqlutil

import (
	"fmt"
	"strings"
)

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

type Condition struct {
	expr strings.Builder
}

func (c *Condition) Append(val string) *Condition {
	fmt.Println("append", val, c.expr.String())

	if c.expr.Len() > 0 {
		c.expr.WriteString(" ")
	}

	c.expr.WriteString(val)
	return c
}

func (c *Condition) Open() *Condition {
	return c.Append("(")
}

func (c *Condition) Close() *Condition {
	return c.Append(")")
}

func (c *Condition) Repeat(delim string, exprs ...string) *Condition {
	if len(exprs) < 1 && delim != "" {
		return c.Append(delim)
	}

	for i, expr := range exprs {
		if i > 0 && delim != "" {
			c.Append(delim)
		}

		c.Append(expr)
	}

	return c
}

func (c *Condition) And(exprs ...string) *Condition {
	return c.Repeat("AND", exprs...)
}

func (c *Condition) Or(exprs ...string) *Condition {
	return c.Repeat("OR", exprs...)
}

func (c *Condition) Like(column, value string) *Condition {
	return c.Append(column).Append("LIKE").Append(value)
}

func (c *Condition) Eq(column, value string) *Condition {
	return c.Append(column).Append("=").Append(value)
}

func (c *Condition) Gt(column, value string) *Condition {
	return c.Append(column).Append(">").Append(value)
}

func (c *Condition) Ge(column, value string) *Condition {
	return c.Append(column).Append(">=").Append(value)
}

func (c *Condition) Lt(column, value string) *Condition {
	return c.Append(column).Append("<").Append(value)
}

func (c *Condition) Le(column, value string) *Condition {
	return c.Append(column).Append("<=").Append(value)
}

func (c *Condition) String() string {
	return c.expr.String()
}
