package main

type rect struct {
	width, height int
}

func (r *rect) calcArea() int {
	return r.width * r.height
}

// func main() {
// 	myrect := rect{width: 10, height: 5}
// 	fmt.Printf("The area is %d\n", myrect.calcArea())
// }
