package main

import (
	"os"
	"testing"
)

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
	t.Run("på norsk bokmål uten navn", func(t *testing.T) {
		got := Hello("", norwegianBokmal)
		want := "Hei, verden"
		assertCorrectMessage(t, got, want)
	})

	t.Run("french bonjour", func(t *testing.T) {
		got := Hello("Jean", french)
		want := "Bonjour, Jean"
		assertCorrectMessage(t, got, want)
	})
	t.Run("french bonjour without name", func(t *testing.T) {
		got := Hello("", french)
		want := "Bonjour, le monde"
		assertCorrectMessage(t, got, want)
	})

	t.Run("spanish hola", func(t *testing.T) {
		got := Hello("José", "es")
		want := "Hola, José"
		assertCorrectMessage(t, got, want)
	})
	t.Run("spanish hola sin nombre", func(t *testing.T) {
		got := Hello("", "es")
		want := "Hola, mundo"
		assertCorrectMessage(t, got, want)
	})
}

func TestParseArg(t *testing.T) {
	t.Run("name and lang", func(t *testing.T) {
		// mess with args, save first https://stackoverflow.com/a/33723649
		originalArgs := os.Args
		defer func() { os.Args = originalArgs }()
		os.Args = []string{"hello.go", "Juan", "es"}
		want := [2]string{"Juan", "es"}
		got := parseArgs()
		assertCorrectArgs(t, got, want)
	})
	t.Run("name only", func(t *testing.T) {
		originalArgs := os.Args
		defer func() { os.Args = originalArgs }()
		os.Args = []string{"hello.go", "Bob"}
		want := [2]string{"Bob", ""}
		got := parseArgs()
		assertCorrectArgs(t, got, want)
	})
	t.Run("empty args", func(t *testing.T) {
		originalArgs := os.Args
		defer func() { os.Args = originalArgs }()
		os.Args = []string{"hello.go"}
		want := [2]string{"", ""}
		got := parseArgs()
		assertCorrectArgs(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // makes the stack trace point to the caller
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectArgs(t testing.TB, got [2]string, want [2]string) {
	t.Helper()
	if len(got) != 2 {
		t.Errorf("got args length %d want %d", len(got), 2)
	}
	assertCorrectMessage(t, got[0], want[0])
	assertCorrectMessage(t, got[1], want[1])
}
