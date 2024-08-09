package main

import "testing"

func TestHelloWorld(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloWorld("Yousef", "")
		want := "Hello, Yousef"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world when empty string provided", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := HelloWorld("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := HelloWorld("Yousef", "French")
		want := "Bonjour, Yousef"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
