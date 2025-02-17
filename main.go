package main

import (
	"log"
	"net/http"

	httpserver "github.com/gppmad/gocloudplayground/http_server"
)

func main() {
	handler := http.HandlerFunc(httpserver.PlayerServer)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
