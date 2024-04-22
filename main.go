package main

import (
	"fmt"

	"github.com/gppmad/gocloudplayground/myutil"
)

func main() {
	// (1) Read from a list, and get list sorted and the max.
	numbers := []int{15, 2, 50, 3, 12}
	orderedNumbers := myutil.GetOrderedNumbers(numbers)
	fmt.Println("The max number is ", orderedNumbers[0])
	fmt.Println("List:")
	for _, num := range orderedNumbers {
		fmt.Println(num)
	}

	// (2) Read a list from the file
	fileNumbers, err := myutil.ReadNumbersFromFile("numbers.txt")
	if err != nil {
		fmt.Println("Error reading from the file:", err)
		return
	}

	fmt.Println("Numbers read from the file:", fileNumbers)

	// TODO Create the first object People and assign to a dictionary.
	// create a function that accepts a dictionary and returns
	// a list of object ordered by age.
	people := make(map[string]int)
	people["peppe"] = 32
	people["turiddu"] = 59
	people["saro"] = 67

	ages := []int{}
	for key, value := range people {
		fmt.Printf("Key %s value %d\n", key, value)
		ages = append(ages, value)
		ages = myutil.GetOrderedNumbers(ages)
	}
	fmt.Println("Ages read:", ages)

}
