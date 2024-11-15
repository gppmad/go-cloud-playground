package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {

	t.Run("with two repetions", func(t *testing.T) {
		got := Repeater("a", 2)
		want := "aa"

		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	})

	t.Run("with five repetions", func(t *testing.T) {
		got := Repeater("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	})

}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 2)
	}
}

func ExampleRepeater() {
	// I want to repeat letter A, 5 times.
	repeatLetter := Repeater("A", 5)
	fmt.Println(repeatLetter)

	// Output: AAAAA
}

func TestRepeaterWord(t *testing.T) {
	want := "peppepeppepeppe"
	got := RepeaterWord("peppe", 3)

	if got != want {
		t.Errorf("got %q expected %q", got, want)
	}
}
