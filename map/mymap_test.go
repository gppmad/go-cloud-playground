package mymap

import (
	"testing"
)

func TestSearch(t *testing.T) {
	x := "uomo di programmazione"

	t.Run("known word", func(t *testing.T) {

		dictionary := Dictionary{"peppe": x}
		got, _ := dictionary.Search("peppe")
		want := "uomo di programmazione"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"peppe": x}
		_, err := dictionary.Search("peppe_and_valeria")

		if err == nil {
			t.Fatal("expected to get an err")
		}

		assertError(t, err, ErrWordNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{"peppe": "uomo di programmazione"}
		dictionary.Add("peppe_vale", "peppe_and_valeria")
		got, err := dictionary.Search("peppe_vale")
		want := "peppe_and_valeria"

		if err != nil {
			t.Fatal("not expected that the value has not been found")
		}

		assertStrings(t, got, want)
		assertDefinition(t, dictionary, "peppe_vale", "peppe_and_valeria")
	})

	t.Run("add another word", func(t *testing.T) {
		dictionary := Dictionary{"peppe": "uomo di programmazione"}
		errAdd := dictionary.Add("peppe", "uomo di programmazione_second_time")

		assertError(t, errAdd, ErrKeyAlreadyExisting)
		assertDefinition(t, dictionary, "peppe", "uomo di programmazione")
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update existing word", func(t *testing.T) {
		dictionary := Dictionary{"peppe": "uomo di programmazione"}
		dictionary.Update("peppe", "peppe_and_valeria")

		assertDefinition(t, dictionary, "peppe", "peppe_and_valeria")
	})

	t.Run("update new word - ?", func(t *testing.T) {
		dictionary := Dictionary{"peppe": "uomo di programmazione"}
		err := dictionary.Update("peppex", "peppe_and_valeria")

		assertError(t, err, ErrWordNotFound)
	})

}

func TestDelete(t *testing.T) {

	t.Run("delete existing word", func(t *testing.T) {
		dictionary := Dictionary{"peppe": "uomo di programmazione"}
		err := dictionary.Delete("peppe")

		assertError(t, err, nil)
	})

	t.Run("delete non existing word - ?", func(t *testing.T) {
		dictionary := Dictionary{"peppe": "uomo di programmazione"}
		err := dictionary.Delete("ciccio")

		assertError(t, err, ErrKeyWorldDoesntExist)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q but want: %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
