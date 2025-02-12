package mycontext

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MyStore struct {
	response string
	t        *testing.T
}

func (s *MyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
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
	case d := <-data:
		return d, nil

	}
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
	})

	t.Run("test cancel", func(t *testing.T) {
		store := &MyStore{response: "Hello", t: t}
		server := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// create request with cancel context
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response should not have been written")
		}

	})

}
