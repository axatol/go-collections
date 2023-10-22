package jsonutil_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/axatol/go-utils/jsonutil"
)

func Test_Duration_MarshalJSON(t *testing.T) {
	tests := []struct {
		input    *jsonutil.Duration
		expected []byte
	}{
		{input: &jsonutil.Duration{time.Duration(0) * time.Millisecond}, expected: []byte("0")},
		{input: &jsonutil.Duration{time.Duration(100) * time.Millisecond}, expected: []byte("100")},
		{input: &jsonutil.Duration{time.Duration(999999999) * time.Millisecond}, expected: []byte("999999999")},
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

func Test_Duration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    []byte
		expected *jsonutil.Duration
	}{
		{input: []byte("0"), expected: &jsonutil.Duration{time.Duration(0) * time.Millisecond}},
		{input: []byte("100"), expected: &jsonutil.Duration{time.Duration(100) * time.Millisecond}},
		{input: []byte("999999999"), expected: &jsonutil.Duration{time.Duration(999999999) * time.Millisecond}},
	}

	for _, sub := range tests {
		sub := sub
		t.Run(string(sub.input), func(t *testing.T) {
			t.Parallel()

			var dur jsonutil.Duration
			err := json.Unmarshal(sub.input, &dur)
			t.Logf("input: %s, expected: %s, actual: %s", sub.input, sub.expected, dur)

			if err != nil {
				t.Fatalf("expected err to be nil, got '%s'", err)
			}

			if sub.expected.Milliseconds() != dur.Milliseconds() {
				t.Fatalf("expected time to be '%s', got '%s'", sub.expected.String(), dur.String())
			}
		})
	}
}
