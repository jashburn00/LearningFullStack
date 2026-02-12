package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Hello should return 'Hello, World!'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		PerformAssertion(t, got, want)
	})

	t.Run("in other languages", func(t *testing.T) {
		got := Hello("Spanish")
		want := "Hola, Mundo!"
		PerformAssertion(t, got, want)
	})
}

func PerformAssertion(t *testing.T, got, want string) {
	t.Helper() //this marks the function as a test helper function, so that when it fails, the line number reported will be in the caller of this function, rather than inside this function itself.
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
