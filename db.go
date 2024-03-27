package main

import (
	"fmt"
	"log"
	"net/http"
)

func anotherHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("printed Another Handler")
	fmt.Fprint(w, "Another Handler!")
}

// SELECT DB Version
