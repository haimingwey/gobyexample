package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people with english", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in chinese", func(t *testing.T) {
		got := Hello("海明威", "chinese")
		want := "你好, 海明威"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Bonjour", "french")
		want := "Hola, Bonjour"
		assertCorrectMessage(t, got, want)
	})
}
