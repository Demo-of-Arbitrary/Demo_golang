package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("calculate any size of numbers' sum", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)
		expected := 6
		if result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("calculate collective array", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4}
		numbers2 := []int{1, 2, 3, 4, 5, 6}
		result := SumAll(numbers1, numbers2)
		expected := [2]int{10, 21}
		if result != expected {
			t.Errorf("expected %+v, got %+v", expected, result)
		}
	})
}
