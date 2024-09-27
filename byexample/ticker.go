// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	done := make(chan bool)
// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				fmt.Println("Received done")
// 				return
// 			case t := <-ticker.C:
// 				fmt.Println("Ticker triggered", t)
// 			}
// 		}

// 	}()
// 	time.Sleep(1600 * time.Millisecond)
// 	done <- true
// 	fmt.Println("Ticker stopped")

// }
