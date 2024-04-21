package myutil

import (
	"reflect"
	"testing"
)

// Test function for getOrderedNumbers
func TestGetOrderedNumbers(t *testing.T) {
	numbers := []int{5, 3, 8, 1, 2}
	expected := []int{8, 5, 3, 2, 1}
	result := GetOrderedNumbers(numbers)

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getOrderedNumbers(%v) = %v, want %v", numbers, result, expected)
	}
}

func TestReadNumbersFromFile(t *testing.T) {
	resultList := []int{15, 2, 50, 3, 12}
	result, err := ReadNumbersFromFile("test_numbers.txt")

	if err != nil {
		t.Fatal("Error reading numbers from file:", err)
	}

	if !reflect.DeepEqual(result, resultList) {
		t.Errorf("expected numbers %v, got %v", resultList, result)
	}
}
