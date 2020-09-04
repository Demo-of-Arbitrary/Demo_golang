package main

import "testing"

func assertString(t *testing.T, want string, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
func assertError(t *testing.T, want, got error) {
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
		want := ErrorNotFound
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("apple", "this is a fruit")
		got, err := dictionary.Search("apple")

		if err != nil {
			t.Fatal("should find added word", err)
		}

		want := "this is a fruit"
		assertString(t, want, got)
	})
}
