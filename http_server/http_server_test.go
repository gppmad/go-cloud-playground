package httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	http_server "github.com/gppmad/gocloudplayground/http_server"
)

func TestGet(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	http_server.PlayerServer(response, request)

	got := response.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
