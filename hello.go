package main

import "fmt"

func Hello(name string) string {
	x := "Hello " + name
	return x
}

func main() {
	x := Hello("peppe")
	fmt.Println(x)
}
