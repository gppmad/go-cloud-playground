package main

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
