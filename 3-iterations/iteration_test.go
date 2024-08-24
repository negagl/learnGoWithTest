package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected %q, but got %q", expected, repeat)
	}
}