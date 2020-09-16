package clockface_test

import (
	"testing"
	"time"

	clockface "github.com/Demo_golang/learn_go_with_tests/maths"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
