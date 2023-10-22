package jsonutil_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/axatol/go-utils/jsonutil"
)

func Test_Time_MarshalJSON(t *testing.T) {
	tests := []struct {
		input    *jsonutil.Time
		expected []byte
	}{
		{input: &jsonutil.Time{time.UnixMilli(0)}, expected: []byte("0")},
		{input: &jsonutil.Time{time.UnixMilli(100)}, expected: []byte("100")},
		{input: &jsonutil.Time{time.UnixMilli(999999999)}, expected: []byte("999999999")},
	}

	for _, sub := range tests {
		sub := sub
		t.Run(string(sub.expected), func(t *testing.T) {
			t.Parallel()

			raw, err := json.Marshal(sub.input)
			t.Logf("input: %s, expected: %s, actual: %s", sub.input, sub.expected, raw)

			if err != nil {
				t.Fatalf("expected err to be nil, got '%s'", err)
			}

			if !bytes.Equal(sub.expected, raw) {
				t.Fatalf("expected raw to be '%s', got '%s'", string(sub.expected), string(raw))
			}
		})
	}
}

func Test_Time_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    []byte
		expected *jsonutil.Time
	}{
		{input: []byte("0"), expected: &jsonutil.Time{time.UnixMilli(0)}},
		{input: []byte("100"), expected: &jsonutil.Time{time.UnixMilli(100)}},
		{input: []byte("999999999"), expected: &jsonutil.Time{time.UnixMilli(999999999)}},
	}

	for _, sub := range tests {
		sub := sub
		t.Run(string(sub.input), func(t *testing.T) {
			t.Parallel()

			var time jsonutil.Time
			err := json.Unmarshal(sub.input, &time)
			t.Logf("input: %s, expected: %s, actual: %s", sub.input, sub.expected, time)

			if err != nil {
				t.Fatalf("expected err to be nil, got '%s'", err)
			}

			if !sub.expected.Equal(time.Time) {
				t.Fatalf("expected time to be '%s', got '%s'", sub.expected.String(), time.String())
			}
		})
	}
}
