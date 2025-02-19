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
		_ = dictionary.Add(addedKey, addedDefinition)

		def, _ := dictionary.Search("car")

		assertStrings(t, def, addedDefinition)
	})

	t.Run("add existing word", func(t *testing.T) {
		err := dictionary.Add("test", "this might fail")

		assertError(t, err, ErrAlreadyExists)
	})

	t.Run("update definition", func(t *testing.T) {
		keyToChange := "test"
		changedDefinition := "new definition"
		_ = dictionary.Update(keyToChange, changedDefinition)

		def, _ := dictionary.Search(keyToChange)

		assertStrings(t, def, changedDefinition)
	})

	t.Run("update definition unknown word", func(t *testing.T) {
		keyToChange := "tes"
		changedDefinition := "new definition"
		err := dictionary.Update(keyToChange, changedDefinition)

		assertError(t, err, ErrNotFound)
	})

	t.Run("delete word", func(t *testing.T) {
		dictionary.Delete("test")

		_, err := dictionary.Search("test")

		assertError(t, err, ErrNotFound)

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
