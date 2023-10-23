package httputil_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/axatol/go-utils/httputil"
	"github.com/stretchr/testify/assert"
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
	// given
	recorder := httptest.NewRecorder()
	response := httputil.Response[map[string]string]{}

	// when
	response.SetStatus(http.StatusOK)
	response.AddHeader("X-Some-Header", "Blah")
	response.SetData(&map[string]string{"foo": "bar"})
	_, err := response.Write(recorder)
	result := recorder.Result()

	// then
	assert.NoError(t, err)
	assert.Equal(t, response.Status, result.StatusCode)
	assert.ObjectsAreEqual(response.Headers, result.Header)
	responseBody, err := json.Marshal(response)
	assert.NoError(t, err)
	resultBody, err := io.ReadAll(result.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, string(responseBody), string(resultBody))
}

func Test_Response_Error(t *testing.T) {
	// given
	recorder := httptest.NewRecorder()
	response := httputil.Response[map[string]string]{}
	expectedErr := fmt.Errorf("blah")

	// when
	response.SetError(expectedErr)
	_, err := response.Write(recorder)
	result := recorder.Result()

	// then
	assert.NoError(t, err)
	assert.Equal(t, response.Status, result.StatusCode)
	responseBody, err := json.Marshal(response)
	assert.NoError(t, err)
	resultBody, err := io.ReadAll(result.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, string(responseBody), string(resultBody))
}
