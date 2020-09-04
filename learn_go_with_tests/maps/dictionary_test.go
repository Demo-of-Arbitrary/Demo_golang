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
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
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
		err := dictionary.Add("apple", "this is a fruit")
		assertError(t, err, nil)
		assertDefinitions(t, dictionary, "apple", "this is a fruit")
	})

	t.Run("should not add word, if the key exists", func(t *testing.T) {
		dictionary := Dictionary{"apple": "this is a fruit"}
		err := dictionary.Add("apple", "this is orange")

		assertError(t, err, ErrWordExists)
		assertDefinitions(t, dictionary, "apple", "this is a fruit")
	})
}

func assertDefinitions(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search("apple")

	if err != nil {
		t.Fatal("should find added word", err)
	}

	want := "this is a fruit"
	assertString(t, want, got)
}
