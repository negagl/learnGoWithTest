package arrayslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("array of 5 items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected %d but got %d, %v", expected, got, numbers)
		}
	})
}
