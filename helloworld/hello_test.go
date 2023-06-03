package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Toine", "")
		want := "Hello, Toine"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, world' when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"

			assertCorrectMessage(t, got, want)
		})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Loic", "French")
		want := "Salut, Loic"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
