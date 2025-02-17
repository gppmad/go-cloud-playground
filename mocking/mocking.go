package oldmain

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord string = "go"
const countDownStart = 3

func Countdown(io io.Writer, s Sleeper) {

	// This print 3,2,1 go!
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(io, i)
		s.Sleep()
	}
	fmt.Fprint(io, finalWord)

}
