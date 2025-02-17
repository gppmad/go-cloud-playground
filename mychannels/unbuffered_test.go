package main

import (
	"testing"
)

func TestJob(t *testing.T) {
	result := job(2)
	expected := 4
	if result != expected {
		t.Errorf("job(2) = %d; want %d", result, expected)
	}
}
