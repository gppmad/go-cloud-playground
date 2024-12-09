package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (ss *SpySleeper) Sleep() {
	ss.Calls++
}

func TestMocking(t *testing.T) {
	buffer := bytes.Buffer{}
	spySleeper := SpySleeper{}

	Countdown(&buffer, &spySleeper)
	got := buffer.String()
	want := `3
2
1
go`

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

}
