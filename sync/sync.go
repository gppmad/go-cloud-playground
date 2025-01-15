package mysync

import "sync"

type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter = c.counter + 1
}

func (c *Counter) Value() int {
	return c.counter
}

func NewCounter() *Counter {
	return &Counter{}
}
