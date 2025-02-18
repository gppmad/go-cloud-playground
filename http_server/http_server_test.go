package httpserver_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	http_server "github.com/gppmad/gocloudplayground/http_server"
)

type StubPlayerStore struct {
	score map[string]int
}

func (ps *StubPlayerStore) GetScore(name string) (int, error) {
	score, ok := ps.score[name]
	if !ok {
		return 0, errors.New("player not found")
	}
	return score, nil
}

func (ps *StubPlayerStore) SetScore(name string) bool {
	if _, ok := ps.score[name]; !ok {
		return false
	}
	ps.score[name]++
	return true
}

func TestGet(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	myserver := &http_server.PlayerServer{Store: &store}

	t.Run("Pepper", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		myserver.ServerHttp(response, request)

		got := response.Body.String()
		want := "20"

		AssertBody(t, got, want)
	})

	t.Run("Floyd", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		myserver.ServerHttp(response, request)

		got := response.Body.String()
		want := "10"

		AssertBody(t, got, want)
	})

	t.Run("Peppe with 400", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Peppe", nil)
		response := httptest.NewRecorder()

		myserver.ServerHttp(response, request)

		got := response.Result().StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestPost(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	myserver := &http_server.PlayerServer{Store: &store}

	t.Run("Save score for Floyd", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		myserver.ServerHttp(response, request)

		got := response.Result().StatusCode
		want := http.StatusAccepted

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}

		floydScore, _ := store.GetScore("Floyd")
		if floydScore != 11 {
			t.Errorf("got %d want %d", floydScore, 11)
		}
	})

}

func AssertBody(t *testing.T, got string, want string) {

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
