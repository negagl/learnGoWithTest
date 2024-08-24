package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nega")
	want := "Hello, Nega!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}