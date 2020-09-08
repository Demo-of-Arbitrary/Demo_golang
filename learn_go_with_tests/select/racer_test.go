package racer

import "testing"

func TestRacer(t *testing.T) {
	t.Run("Racer", func(t *testing.T) {
		const slowUrl = "http://www.facebook.com"
		const fastUrl = "http://www.quii.co.uk"

		want := fastUrl
		got := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
