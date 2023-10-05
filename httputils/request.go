package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response[T any] struct {
	Headers map[string]string `json:"-"`
	Status  int               `json:"-"`
	Error   error             `json:"-"`
	Message string            `json:"message"`
	Data    *T                `json:"data"`
}

func (r *Response[T]) AddHeader(key, value string) *Response[T] {
	r.Headers[key] = value
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
		}

		r.Status = http.StatusOK
	}

	if r.Message == "" {
		if r.Error != nil {
			r.Message = r.Error.Error()
		} else if r.Status >= 200 && r.Status < 300 {
			r.Message = "OK"
		} else if r.Status >= 400 && r.Status < 500 {
			r.Message = "Bad request"
		} else if r.Status >= 500 {
			r.Message = "An error occurred"
		}
	}

	raw, err := json.Marshal(w)
	if err != nil {
		return 0, fmt.Errorf("failed to serialise response body: %s", err)
	}

	for k, v := range r.Headers {
		w.Header().Add(k, v)
	}

	w.WriteHeader(r.Status)
	return w.Write(raw)
}
