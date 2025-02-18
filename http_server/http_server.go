package httpserver

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetScore(name string) (int, error)
	SetScore(name string) bool
}

type PlayerServer struct {
	Store PlayerStore
}

func (ps *PlayerServer) ServerHttp(response http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodGet:
		ps.GetScore(response, request)
	case http.MethodPost:
		ps.SetScore(response, request)
	}

}

func (ps *PlayerServer) GetScore(response http.ResponseWriter, request *http.Request) {

	player := strings.TrimPrefix(request.URL.Path, "/players/")
	score, err := ps.Store.GetScore(player)
	if err != nil {
		http.Error(response, "Player not found", http.StatusBadRequest)
		return
	}

	fmt.Fprint(response, score)

}

func (ps *PlayerServer) SetScore(response http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	success := ps.Store.SetScore(player)
	if !success {
		http.Error(response, "Unable to set score", http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusAccepted)
}
