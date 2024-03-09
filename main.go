package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("printed Hello World")
	fmt.Fprint(w, "Hello, World!")

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/another-handler", anotherHandler)
	log.Println("Server starting on port 10000")
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatal("Error starting the web server: ", err)
	}

}
