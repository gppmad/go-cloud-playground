// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Container struct {
// 	mu       sync.Mutex
// 	counters map[string]int
// }

// func (c *Container) add(el string) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.counters[el]++
// }

// func main() {
// 	c := Container{
// 		counters: map[string]int{"a": 0, "b": 0},
// 	}

// 	var wg sync.WaitGroup

// 	doIncrement := func(name string, value int) {
// 		for i := 0; i < value; i++ {
// 			c.add(name)
// 		}
// 		wg.Done()
// 	}

// 	wg.Add(2)
// 	go doIncrement("a", 1000)
// 	go doIncrement("b", 1000)

// 	wg.Wait()
// 	fmt.Println(c)
// }
