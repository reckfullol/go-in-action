package chapter_14

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//func (s *SpyStore) assertWasCancelled() {
//	s.t.Helper()
//	if !s.cancelled {
//		s.t.Errorf("store was not told to cancel")
//	}
//}
//
//func (s *SpyStore) assertWasNotCancelled() {
//	s.t.Helper()
//	if s.cancelled {
//		s.t.Errorf("store was told to cancel")
//	}
//}

func TestHandler(t *testing.T) {
	data := "hello, world"
	//svr := Server(&StubStore{data})
	//
	//request := httptest.NewRequest(http.MethodGet, "/", nil)
	//response := httptest.NewRecorder()
	//
	//svr.ServeHTTP(response, request)
	//
	//if response.Body.String() != data {
	//	t.Errorf("got %s, want %s", response.Body.String(), data)
	//}

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Canceling this context releases resources associated with it
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		//store.assertWasCancelled()
	})

	t.Run("return data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, want %s", response.Body.String(), data)
		}

		//store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not have been written")
		}
	})
}
