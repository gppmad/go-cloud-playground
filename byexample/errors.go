package main

import (
	"errors"
	"fmt"
)

func check(arg int) (int, error) {
	if arg == 10 {
		return -1, errors.New("i cannot work with 10")
	}

	return arg, nil
}

var BasicError = fmt.Errorf("basic error")

func raiseError(arg int) error {
	if arg == 2 {
		return fmt.Errorf("another error on top of %w", BasicError)
	}
	return fmt.Errorf("another error")
}

// func main() {

// 	for i := 0; i <= 10; i++ {
// 		if r, e := check(i); e != nil {
// 			fmt.Println("The check doesn't pass with error:", e)
// 		} else {
// 			fmt.Println("The check passes", r)
// 		}
// 	}

// 	err := raiseError(4)
// 	if errors.Is(err, BasicError) {
// 		fmt.Println("Returned a basic error")
// 	} else {
// 		fmt.Println(err)
// 	}
// }
