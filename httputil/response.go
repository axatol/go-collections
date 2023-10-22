package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response[T any] struct {
	Headers http.Header `json:"-"`
	Status  int         `json:"-"`
	Error   error       `json:"-"`
	Message string      `json:"message"`
	Data    *T          `json:"data"`
}

func (r *Response[T]) AddHeader(key, value string) *Response[T] {
	if r.Headers == nil {
		r.Headers = http.Header{}
	}

	r.Headers[key] = append(r.Headers[key], value)
	return r
}

func (r *Response[T]) SetStatus(status int) *Response[T] {
	r.Status = status
	return r
}

func (r *Response[T]) SetMessage(message string) *Response[T] {
	r.Message = message
	return r
}

func (r *Response[T]) SetError(err error) *Response[T] {
	r.Error = err
	return r
}

func (r *Response[T]) SetData(data *T) *Response[T] {
	r.Data = data
	return r
}

func (r *Response[T]) Write(w http.ResponseWriter) (int, error) {
	if r.Status == 0 {
		if r.Error != nil {
			r.Status = http.StatusInternalServerError
		} else {
			r.Status = http.StatusOK
		}
	}

	if r.Message == "" {
		if r.Error != nil {
			r.Message = r.Error.Error()
		} else {
			r.Message = http.StatusText(r.Status)
		}
	}

	raw, err := json.Marshal(r)
	if err != nil {
		return 0, fmt.Errorf("failed to serialise response body: %s", err)
	}

	for name, values := range r.Headers {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(r.Status)
	return w.Write(raw)
}
