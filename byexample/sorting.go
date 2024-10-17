/* package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Giuseppe", age: 32},
		Person{name: "Salvatore", age: 60},
		Person{name: "Mauro", age: 28},
	}

	compareFunc := func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	}

	slices.SortFunc(people, compareFunc)
	fmt.Println(people)
}
*/