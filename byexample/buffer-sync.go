// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(done chan bool) {
// 	fmt.Println("Doing something")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- true
// }

// /* func main() {
// 	doneChannel := make(chan bool, 1)
// 	go worker(doneChannel)
// 	//<-doneChannel

// } */
