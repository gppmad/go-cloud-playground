package iteration

import "strings"

// Create a word with a letter repeated N times.
func Repeater(letter string, times int) string {
	var word string
	for i := 0; i < times; i++ {
		word += letter
	}
	return word
}

func RepeaterWord(word string, times int) string {
	var finalWord string
	for i := 0; i < times; i++ {
		finalWord += strings.Clone(word)
	}
	return finalWord
}
