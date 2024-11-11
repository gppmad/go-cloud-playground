package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello peppe"
	got := Hello("peppe")

	if got != expected {
		t.Errorf("Got %q instead of %q", got, expected)
	}
}
