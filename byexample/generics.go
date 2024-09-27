package main

import "fmt"

// Let's playing around with linked list

type List[T any] struct {
	head, tail *Element[T]
}

type Element[T any] struct {
	next *Element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	el := &Element[T]{val: v}
	if lst.tail == nil {
		lst.head = el
		lst.tail = lst.head
	} else {
		lst.tail.next = el
		lst.tail = lst.head.next
	}
}

func (lst *List[T]) GetAll() []T {
	var slc []T
	for e := lst.head; e != nil; e = e.next {
		slc = append(slc, e.val)
	}
	return slc
}

func SliceIndex[S []E, E comparable](s S, e E) int {
	for i := range s {
		if s[i] == e {
			return i
		}
	}
	return -1
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// func main() {
// 	// intElements := []int{1, 2, 4}
// 	// stringElements := []string{"el1", "el2", "el3"}

// 	// // I want to find out which position the turkey is in
// 	// animals := []string{"monkey", "turkey", "dog"}

// 	myList := List[int]{}
// 	myList.Push(2)
// 	fmt.Println(myList.GetAll())

// }
