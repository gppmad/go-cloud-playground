package arrays

import "testing"

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
