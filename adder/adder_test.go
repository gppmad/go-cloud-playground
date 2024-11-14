package adder

import (
	"fmt"
	"testing"
)

// TestAdd checks that the Add function works as expected.
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("Add(2, 3) = %v; want %v", got, want)
	}
}

func TestAddNegative(t *testing.T) {
	got := Add(-1, -1)
	want := -2

	if got != want {
		t.Errorf("Add(-1, -1) = %v; want %v", got, want)
	}
}

func ExampleAdd() {
	got := Add(3, 5)
	fmt.Println(got)

	// Output: 8
}
