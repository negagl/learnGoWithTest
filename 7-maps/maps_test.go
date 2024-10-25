package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("Was expecting an error...")
		}

		assertError(t, got, ErrNotFound)
	})

	t.Run("add word", func(t *testing.T) {
		addedKey := "car"
		addedDefinition := "vehicle to transport"
		dictionary.Add(addedKey, addedDefinition)

		def, _ := dictionary.Search("car")

		assertStrings(t, def, addedDefinition)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got %q wanted %q", got, want)
	}
}
