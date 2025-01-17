package mycontext

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MyStore struct {
	response string
	cancel   bool
	t        *testing.T
}

func (s *MyStore) Fetch() string {
	time.Sleep(10 * time.Millisecond)
	return s.response
}

func (s *MyStore) Cancel() {
	s.cancel = true
}

func (s *MyStore) AssertWasCancel() {
	s.t.Helper()
	if !s.cancel {
		s.t.Errorf("the store has not been cancelled")
	}
}

func (s *MyStore) AssertWasNotCancel() {
	s.t.Helper()
	if s.cancel {
		s.t.Errorf("the store has been cancelled")
	}
}

func TestServer(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		output := "Hello"
		store := &MyStore{response: "Hello", t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != output {
			t.Errorf("got %s, want %s", response.Body.String(), output)
		}

		store.AssertWasNotCancel()
	})

	t.Run("test cancel", func(t *testing.T) {
		store := &MyStore{response: "Hello", t: t}
		server := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// create request with cancel context
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		store.AssertWasCancel()

	})

}
