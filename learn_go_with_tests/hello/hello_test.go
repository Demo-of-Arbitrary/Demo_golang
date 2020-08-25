package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("George", "")
		want := "Hello George"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World', if no name passed in", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola', if passed lang=spanish", func(t *testing.T) {
		got := Hello("world", "spanish")
		want := "Hola world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Bonjour', if passed lang=french", func(t *testing.T) {
		got := Hello("world", "french")
		want := "Bonjour world"
		assertCorrectMessage(t, got, want)
	})
}
