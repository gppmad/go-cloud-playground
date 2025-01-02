package myselect

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(timeout time.Duration, a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

}

func ping(url string) chan struct{} {
	res := make(chan struct{})

	go func() {
		http.Get(url)
		close(res)
	}()

	return res
}
