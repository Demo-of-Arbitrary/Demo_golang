package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})
}

func assertCounter(t *testing.T, counter Counter, got int) {
	want := counter.Value()
	t.Helper()
	if want != 3 {
		t.Errorf("got %d, expected %d", want, got)
	}
}
