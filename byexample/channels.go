package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "coming from the goroutine" }()

	msg := <-messages
	fmt.Println(msg)
}
