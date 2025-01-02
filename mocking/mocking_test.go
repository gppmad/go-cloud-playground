package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyOperations struct {
	Calls []string
}

func (s *SpyOperations) Write([]byte) (int, error) {
	s.Calls = append(s.Calls, "write")
	return 0, nil
}

func (s *SpyOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func TestMocking(t *testing.T) {

	t.Run("high-level behavior", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := SpyOperations{}

		Countdown(&buffer, &spySleeper)
		got := buffer.String()
		want := `3
2
1
go`

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}

	})

	t.Run("correct sequence", func(t *testing.T) {
		spyOP := SpyOperations{}
		Countdown(&spyOP, &spyOP)
		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(spyOP.Calls, want) {
			t.Fatalf("got: %q, want %q", spyOP.Calls, want)
		}

	})

}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

type SpySleeper struct {
	durationSlept time.Duration
}

func (s *SpySleeper) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestSleeper(t *testing.T) {
	spySleeper := SpySleeper{durationSlept: 5}
	cf := ConfigurableSleeper{5, spySleeper.Sleep}
	cf.Sleep()

	if spySleeper.durationSlept != 5 {
		t.Fatalf("got %v, want %v", cf.duration, 5)
	}
}
