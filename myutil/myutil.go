package myutil

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetOrderedNumbers(numbers []int) []int {

	orderedNumbers := numbers
	sort.Slice(orderedNumbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	return orderedNumbers
}

func ReadNumbersFromFile(fileName string) ([]int, error) {

	// Open the text file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	// Slice to store the numbers
	var numbers []int

	// Read line by line and convert each line to a number
	for scanner.Scan() {
		// Convert the text to an integer
		line := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting line to integer:", err)
			return nil, err
		}
		// Append the number to the slice
		numbers = append(numbers, num)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return nil, err
	}

	// Print the numbers read from the file
	return numbers, nil
}
