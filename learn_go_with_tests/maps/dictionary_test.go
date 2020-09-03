package main

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Dictionary", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}
		got := Search(dictionary, "test")
		want := "this is just a test"

		if got != want {
			t.Errorf("want %q, got %q, given %q", want, got, "test")
		}
	})
}
