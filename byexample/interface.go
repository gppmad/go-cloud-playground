package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

// func main() {

// 	r := rectangle{width: 5, height: 2}
// 	c := circle{radius: 1}
// 	measure(r)
// 	measure(c)
// }
