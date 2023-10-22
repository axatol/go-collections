package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	_ http.ResponseWriter = (*MockResponseWriter)(nil)
)

type MockResponseWriter struct {
	Headers   http.Header
	WriteHook func([]byte) (int, error)

	WrittenStatusCode int
	WrittenHeaders    http.Header
	WrittenBody       []byte
}

func (w *MockResponseWriter) Header() http.Header {
	if w.Headers == nil {
		w.Headers = http.Header{}
	}

	return w.Headers
}

func (w *MockResponseWriter) Write(input []byte) (int, error) {
	w.WrittenHeaders = w.Headers.Clone()
	w.WrittenBody = input

	if w.WriteHook != nil {
		return w.WriteHook(input)
	}

	return len(input), nil
}

func (w *MockResponseWriter) WriteHeader(statusCode int) {
	w.WrittenStatusCode = statusCode
}

func (w *MockResponseWriter) TestEqualsHeaders(headers http.Header) error {
	if len(headers) != len(w.WrittenHeaders) {
		return fmt.Errorf("expected %d headers, got %d", len(headers), len(w.WrittenHeaders))
	}

	for name, expected := range headers {
		actual, ok := w.WrittenHeaders[name]
		if !ok {
			return fmt.Errorf("expected header %s to exit, but it did not", name)
		}

		if len(expected) != len(actual) {
			return fmt.Errorf("expected header %s to have value of length %d, got %d", name, len(expected), len(actual))
		}

		for i, val := range expected {
			if val != actual[i] {
				return fmt.Errorf("expected header %s[%d] to be '%s', got '%s'", name, i, val, actual[i])
			}
		}
	}

	return nil
}

func (w *MockResponseWriter) TestEqualsBody(expected any) error {
	expectedRaw, err := json.Marshal(expected)
	if err != nil {
		return fmt.Errorf("failed to marshal expected value: %s", err)
	}

	if !bytes.Equal(expectedRaw, w.WrittenBody) {
		return fmt.Errorf("expected body to be '%s', got '%s'", string(expectedRaw), string(w.WrittenBody))
	}

	return nil
}
