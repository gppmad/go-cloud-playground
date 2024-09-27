package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("My internal number is %d", b.num)
}

type container struct {
	base
	s string
}

type element interface {
	describe() string
}

// func main() {
// 	b := base{num: 1}
// 	c := container{base: b, s: "hello"}

// 	fmt.Printf("My container num is: %d, and the string is %v\n", c.num, c.s)
// 	fmt.Println(c.describe())

// 	var el element = c
// 	fmt.Println(el.describe())
// }
