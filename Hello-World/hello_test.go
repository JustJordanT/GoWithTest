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
		got := Hello("JT")
		want := "Hello, JT"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, world' when a empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}
