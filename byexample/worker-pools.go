// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, result chan<- int, wg *sync.WaitGroup) {
// 	fmt.Println("Starting Worker", id)
// 	result <- id * 2
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// jobs := make(chan int)
// 	result := make(chan int, 5) // Buffered channel with capacity 5

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		worker(i, result, &wg)
// 	}

// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Result:", <-result)
// 	}

// 	wg.Wait()
// }
