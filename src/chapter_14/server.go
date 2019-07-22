package chapter_14

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//// Return a channel which gets sent a signal when the context is done or cancelled
		//ctx := request.Context()
		//data := make(chan string, 1)
		//
		//go func() {
		//	data <- store.Fetch()
		//}()
		//
		//select {
		//case d := <- data:
		//	_, _ = fmt.Fprint(writer, d)
		//case <-ctx.Done():
		//	store.Cancel()
		//}

		data, err := store.Fetch(request.Context())

		if err != nil {
			return // todo: log error however you like
		}
		_, _ = fmt.Fprint(writer, data)
	}
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

type SpyStore struct {
	response string
	//cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	//s.cancelled = true
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
