package main

import (
	"fmt"
)

func main() {
	// people := make(map[string]myutil.Person)
	// people["peppe"] = myutil.Person{Name: "Giuseppe", Surname: "Rossi", Age: 32}
	// people["turiddu"] = myutil.Person{Name: "Salvatore", Surname: "Verdi", Age: 59}
	// people["saro"] = myutil.Person{Name: "Rosario", Surname: "Bianchi", Age: 67}

	// orderedPeople := myutil.GetOrderedPeople(people)
	// for _, person := range orderedPeople {
	// 	fmt.Printf("Name: %s, Surname: %s, Age: %d\n", person.Name, person.Surname, person.Age)
	testString := "Hello, world!"
	fmt.Println("With PrintLN")
	for _, char := range testString {
		fmt.Println(char)
	}
	fmt.Println("With Printf")
	for _, char := range testString {
		fmt.Printf("%c\n", char)
	}
}
