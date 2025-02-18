package main

import (
	"log"
	"net/http"

	httpserver "github.com/gppmad/gocloudplayground/http_server"
)

type MyPlayerStore struct {
}

func (ps *MyPlayerStore) GetScore(name string) (int, error) {
	if name == "Pepper" {
		return 20, nil
	}

	if name == "Floyd" {
		return 10, nil
	}
	return 0, nil
}

func (ps *MyPlayerStore) SetScore(name string) bool {
	return true
}

func main() {
	store := MyPlayerStore{}
	server := httpserver.PlayerServer{Store: &store}
	handler := http.HandlerFunc(server.ServerHttp)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
