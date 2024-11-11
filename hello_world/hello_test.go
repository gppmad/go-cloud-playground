package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, peppe"
		got := Hello("peppe", "English")
		assertCorrectMessage(t, got, expected)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		expected := "Hello, World"
		got := Hello("", "English")

		assertCorrectMessage(t, got, expected)
	})

	t.Run("spanish", func(t *testing.T) {
		expected := "Hola, Peppe"
		got := Hello("Peppe", "Spanish")

		assertCorrectMessage(t, got, expected)
	})

	t.Run("french", func(t *testing.T) {
		expected := "Bonjour, Peppe"
		got := Hello("Peppe", "French")

		assertCorrectMessage(t, got, expected)
	})

}

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("Got %q instead of %q", got, expected)
	}
}
