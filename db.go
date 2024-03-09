package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO: Transform this function in something that fetch data from db.
func anotherHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("printed Another Handler")
	fmt.Fprint(w, "Another Handler!")
}

func simpleSum(a int, b int) int {
	return a + b
}
