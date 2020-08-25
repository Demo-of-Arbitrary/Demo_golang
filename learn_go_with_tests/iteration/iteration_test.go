package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		actual := Iteration("k")
		expected := "kkkkk"
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})
}
