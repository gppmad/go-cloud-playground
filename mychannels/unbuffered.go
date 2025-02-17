package main

import (
	"fmt"
	"time"
)

func job(el int) int {
	// this worker is responsible to double a number
	return el * 2
}

func nomain() {
	syncChannel := make(chan int)

	for i := 1; i <= 10; i++ {
		go func(i int) {
			fmt.Println("co routine number", i)
			time.Sleep(5 * time.Second)
			syncChannel <- job(2)
		}(i)
	}

	select {
	case res := <-syncChannel:
		fmt.Println("one of the coroutines has been finished ", res)

	case <-time.After(2 * time.Second):
		fmt.Println("I am impatient, I cannot wait")
	}
}
