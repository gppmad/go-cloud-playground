package main

import "fmt"

func mayPanic() {
	panic("Unexpected Error")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	mayPanic()
}
