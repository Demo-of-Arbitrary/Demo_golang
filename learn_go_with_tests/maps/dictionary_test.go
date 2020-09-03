package main

import "testing"

func assertString(t *testing.T, want string, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestSearch(t *testing.T) {
	t.Run("Dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, want, got)
	})
}
