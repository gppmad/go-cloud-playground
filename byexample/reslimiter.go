// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	requests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		requests <- i
// 	}
// 	close(requests)

// 	limiter := time.Tick(5 * time.Second)
// 	burstyLimiter := make(chan time.Time, 3)
// 	for i := 0; i < 3; i++ {
// 		burstyLimiter <- time.Now()
// 	}
// 	go func() {
// 		for t := range time.Tick(200 * time.Millisecond) {
// 			burstyLimiter <- t
// 		}
// 	}()

// 	for el := range requests {
// 		fmt.Println("processing", el)
// 		<-limiter
// 		fmt.Println(el, time.Now())
// 	}

// }
