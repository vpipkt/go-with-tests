package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("Jason", "")
		want := "Hello, Jason"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello to everyone", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("på norsk bokmål", func(t *testing.T) {
		got := Hello("Idrun", norwegianBokmal)
		want := "Hei, Idrun"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // makes the stack trace point to the caller
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
