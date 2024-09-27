// package main

// import "fmt"

// func worker(queue chan<- string) {
// 	queue <- "Task1"
// 	queue <- "Task2"
// 	close(queue)
// }

// func main() {
// 	queue := make(chan string, 5)
// 	go worker(queue)

// 	for el := range queue {
// 		fmt.Println(el)
// 	}
// }
