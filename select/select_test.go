package myselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare two urls", func(t *testing.T) {
		slowServer := createTestWebServer(20 * time.Millisecond)
		fastServer := createTestWebServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		comparison, _ := Racer(10*time.Second, slowUrl, fastUrl)

		if comparison != fastUrl {
			t.Errorf("got %q, want %q", comparison, fastUrl)
		}
	})

	t.Run("timeout in case one url takes more than 10s", func(t *testing.T) {
		slowServer := createTestWebServer(6 * time.Millisecond)
		superSlowServer := createTestWebServer(10 * time.Millisecond)

		defer slowServer.Close()
		defer superSlowServer.Close()

		_, err := Racer(5*time.Millisecond, slowServer.URL, superSlowServer.URL)

		if err == nil {
			t.Errorf("one of the websites took more than 10s, expected an error butr didn't get one")
		}
	})

}

func createTestWebServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
