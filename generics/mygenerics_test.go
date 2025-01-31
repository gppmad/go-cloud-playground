package mygenerics_test

import (
	"testing"

	mygenerics "github.com/gppmad/gocloudplayground/generics"
)

func assertTrue(t *testing.T, flag bool) {
	t.Helper()
	if flag != true {
		t.Errorf("the condition is not True")
	}
}

func assertFalse(t *testing.T, flag bool) {
	t.Helper()
	if flag != false {
		t.Errorf("the condition is not False")
	}
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v but want %+v", got, want)
	}
}

func TestStack(t *testing.T) {
	myStack := new(mygenerics.MyStack[int])

	assertTrue(t, myStack.IsEmpty())

	myStack.Put(10)
	assertFalse(t, myStack.IsEmpty())

	el, _ := myStack.Pop()
	assertEqual(t, el, 10)

	el, flag := myStack.Pop()
	assertEqual(t, el, 0)
	assertFalse(t, flag)

}
