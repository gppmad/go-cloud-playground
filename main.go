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

	// (2) Compare two lists and print ok if are the same
	fileNumbers, err := myutil.ReadNumbersFromFile("numbers.txt")
	if err != nil {
		fmt.Println("Error reading from the file:", err)
		return
	}

	fmt.Println("Numbers read from the file:", fileNumbers)

}
