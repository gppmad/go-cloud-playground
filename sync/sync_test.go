package mysync

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("The counter has not been called %d Times but %d", want, got.Value())
	}
}

func TestCounter(t *testing.T) {
	c := &Counter{}
	c.Incr()
	c.Incr()
	c.Incr()

	assertCounter(t, c, 3)
}

func TestCounterConcurrently(t *testing.T) {
	n := 1000
	c := NewCounter()

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			c.Incr()
			wg.Done()
		}()
	}
	wg.Wait()
	assertCounter(t, c, n)
}
