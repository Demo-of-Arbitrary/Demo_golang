package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		actual := Repeat("k", 5)
		expected := "kkkkk"
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("k", 5))
	// Output: kkkkk
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000)
	}
}
