package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("Nega", "")
		want := "Hello, Nega!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Nega", "ES")
		want := "Hola, Nega!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Korean", func(t *testing.T) {
		got := Hello("", "KR")
		want := "World 안녕하세요!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}