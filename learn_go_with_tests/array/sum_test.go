package arrays

import (
	"reflect"
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

func TestSumAllTails(t *testing.T) {
	t.Run("calculate collective array", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4}
		numbers2 := []int{1, 2, 3, 4, 5, 6}
		numbers3 := []int{9, 100}
		result := SumAllTails(numbers1, numbers2, numbers3)
		expected := []int{9, 20, 100}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %+v, got %+v", expected, result)
		}
	})
}
