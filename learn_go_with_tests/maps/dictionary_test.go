package main

import "testing"

func assertString(t *testing.T, want string, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, want, got)
	})

	t.Run("unknown Word", func(t *testing.T) {
		_, err := dictionary.Search("good")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertString(t, err.Error(), want)
	})
}
