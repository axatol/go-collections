package httputil_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/axatol/go-utils/httputil"
)

func Test_Response_AddHeader(t *testing.T) {
	headers := map[string]string{
		"foo":   "bar",
		"lorem": "ipsum",
	}

	r := httputil.Response[string]{}
	for key, value := range headers {
		r.AddHeader(key, value)
	}

	for key, expected := range headers {
		actual, ok := r.Headers[key]
		if !ok {
			t.Fatalf("expected header '%s' to exist", key)
		}

		if len(actual) != 1 || expected != actual[0] {
			t.Fatalf("expected header '%s' to have value '%s', got: '%s'", key, expected, actual)
		}
	}
}

func Test_Response_Write(t *testing.T) {
	w := httputil.MockResponseWriter{}
	r := httputil.Response[map[string]string]{}
	r.SetStatus(http.StatusOK)
	r.AddHeader("X-Some-Header", "Blah")
	r.SetData(&map[string]string{"foo": "bar"})

	if _, err := r.Write(&w); err != nil {
		t.Fatalf("expected err to be nil, got: %s", err)
	}

	if r.Status != w.WrittenStatusCode {
		t.Fatalf("expected status code to be %d, got %d", r.Status, w.WrittenStatusCode)
	}

	if err := w.TestEqualsHeaders(r.Headers); err != nil {
		t.Fatal(err)
	}

	if err := w.TestEqualsBody(r); err != nil {
		t.Fatal(err)
	}
}

func Test_Response_Error(t *testing.T) {
	w := httputil.MockResponseWriter{}
	r := httputil.Response[map[string]string]{}
	expectedErr := fmt.Errorf("blah")
	r.SetError(expectedErr)

	if _, err := r.Write(&w); err != nil {
		t.Fatalf("expected err to be nil, got: %s", err)
	}

	if r.Status != w.WrittenStatusCode {
		t.Fatalf("expected status code to be %d, got %d", r.Status, w.WrittenStatusCode)
	}

	if err := w.TestEqualsBody(r); err != nil {
		t.Fatal(err)
	}
}
