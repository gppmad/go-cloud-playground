// db_test.go
package main

import (
	"testing"
)

func TestSimpleSum(t *testing.T) {
	// Test case 1
	result := simpleSum(3, 7)
	expected := 10
	if result != expected {
		t.Errorf("simpleSum(3, 7) = %d; expected %d", result, expected)
	}

	// Test case 2
	result = simpleSum(-5, 8)
	expected = 3
	if result != expected {
		t.Errorf("simpleSum(-5, 8) = %d; expected %d", result, expected)
	}

	// Add more test cases as needed
}
