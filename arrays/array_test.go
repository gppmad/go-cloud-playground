package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of N", func(t *testing.T) {
		want := 31
		el := []int{1, 3, 1, 5, 10, 11}
		got := Sum(el)

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, el)
		}
	})
}

func TestSumAll(t *testing.T) {
	want := []int{3, 9}
	got := SumAll([]int{1, 2}, []int{3, 6})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Sum of some elements", func(t *testing.T) {
		want := []int{1, 8}
		got := SumTails([]int{5, 1}, []int{3, 5, 3})

		checkSums(t, got, want)
	})

	t.Run("Sum of empty elements", func(t *testing.T) {
		want := []int{0, 8}
		got := SumTails([]int{}, []int{3, 5, 3})
		checkSums(t, got, want)
	})

}
