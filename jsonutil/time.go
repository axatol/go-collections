package jsonutil

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Time)(nil)
	_ json.Unmarshaler = (*Time)(nil)
)

// a time that marshals and unmarshals into integer unix milisecond representation
type Time struct{ time.Time }

func (t *Time) MarshalJSON() ([]byte, error) {
	milli := t.UnixMilli()
	str := strconv.FormatInt(milli, 10)
	return []byte(str), nil
}

func (t *Time) UnmarshalJSON(raw []byte) error {
	milli, err := strconv.ParseInt(string(raw), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.UnixMilli(milli)
	return nil
}
