package main

import (
	"fmt"
)

func main() {
	// (1) Read from a list, and get list sorted and the max.
	numbers := []int{15, 2, 50, 3, 12}
	orderedNumbers := GetOrderedNumbers(numbers)
	fmt.Println("The max number is ", orderedNumbers[0])
	fmt.Println("List:")
	for _, num := range orderedNumbers {
		fmt.Println(num)
	}

	// (2) Compare two lists and print ok if are the same
	// 	   find out time complexity for the basics operations with lists.

}
