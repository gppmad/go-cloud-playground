// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type ReadOp struct {
// 	key  int
// 	resp chan float32
// }

// type WriteOp struct {
// 	key   int
// 	value float32
// 	resp  chan bool
// }

// func main() {

// 	reads := make(chan ReadOp)
// 	writes := make(chan WriteOp)

// 	go func() {
// 		// key number: value doubled number
// 		numbers := make(map[int]float32)
// 		numbers[0] = 0
// 		numbers[1] = 1
// 		numbers[2] = 4
// 		numbers[3] = 6

// 		for {
// 			select {
// 			case read := <-reads:
// 				if value, ok := numbers[read.key]; ok {
// 					read.resp <- value
// 				} else {
// 					read.resp <- -1
// 				}

// 			case write := <-writes:
// 				if _, ok := numbers[write.key]; ok {
// 					// I want to preserve existing values
// 					write.resp <- false
// 				} else {
// 					numbers[write.key] = write.value
// 					write.resp <- true
// 				}
// 			}
// 		}
// 	}()

// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			op := ReadOp{key: i, resp: make(chan float32)}
// 			reads <- op
// 			resp := <-op.resp
// 			fmt.Println("Requested", op.key, "got", resp)
// 			wg.Done()
// 		}()
// 	}

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			op := WriteOp{key: i, value: float32(i * 2), resp: make(chan bool)}
// 			writes <- op
// 			resp := <-op.resp
// 			fmt.Println("Requested to write", op.key, "got", resp)
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// }
