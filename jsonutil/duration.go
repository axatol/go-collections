package jsonutil

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	_ json.Marshaler   = (*Duration)(nil)
	_ json.Unmarshaler = (*Duration)(nil)
)

// a duration that marshals and unmarshals into integer unix milisecond representation
type Duration struct{ time.Duration }

func (d *Duration) MarshalJSON() ([]byte, error) {
	milli := d.Milliseconds()
	str := strconv.FormatInt(milli, 10)
	return []byte(str), nil
}

func (d *Duration) UnmarshalJSON(raw []byte) error {
	milli, err := strconv.ParseInt(string(raw), 10, 64)
	if err != nil {
		return err
	}

	d.Duration = time.Duration(milli) * time.Millisecond
	return nil
}
