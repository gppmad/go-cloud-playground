package main

import "sort"

func GetOrderedNumbers(numbers []int) []int {

	orderedNumbers := numbers
	sort.Slice(orderedNumbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	return orderedNumbers
}
