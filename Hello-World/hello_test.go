package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello", func(t *testing.T) {
		got := Hello("JT", "English")
		want := "Hello, JT"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, world' when a empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Josue", "Spanish")
		want := "Hola, Josue"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Philip", "French")
		want := "Bonjour, Philip"
		assertCorrectMessage(t, got, want)
	})
}
