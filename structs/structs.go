package structs

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func (tr Triangle) Area() float64 {
	return tr.Base * tr.Height * 0.5
}
