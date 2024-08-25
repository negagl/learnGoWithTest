package arrayslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array of 5 items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d but got %d, %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3, 4}, []int{0, 9}, []int{})
	want := []int{10, 9, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but got %v", want, got)
		}
	}
	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9, 3, 5})
		want := []int{2, 17}
	
		checkSums(t, got, want)
	})
	t.Run("sum all tails with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
	
		checkSums(t, got, want)
	})
}