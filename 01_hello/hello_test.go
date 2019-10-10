package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Error("got $q want $q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
