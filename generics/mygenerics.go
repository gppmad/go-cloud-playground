package mygenerics

type MyStack[T any] struct {
	values []T
}

func (ms *MyStack[T]) IsEmpty() bool {
	return len(ms.values) == 0
}

func (ms *MyStack[T]) Put(el T) {
	ms.values = append(ms.values, el)
}

func (ms *MyStack[T]) Pop() (T, bool) {
	if len(ms.values) == 0 {
		var zero T
		return zero, false
	}

	index := len(ms.values) - 1
	el := ms.values[index]
	ms.values = ms.values[:index]
	return el, true
}
