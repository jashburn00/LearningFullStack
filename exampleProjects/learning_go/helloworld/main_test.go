package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Hello should return 'Hello, World!'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		PerformStringAssertion(t, got, want)
	})

	t.Run("in other languages", func(t *testing.T) {
		got := Hello("Spanish")
		want := "Hola, Mundo!"
		PerformStringAssertion(t, got, want)
	})
}

func TestWordCount(t *testing.T) {
	t.Run("counts the occurrence of each word in a string", func(t *testing.T) {
		got := WordCount("hello world and love thine enemy for this is the word of the Lord")
		want := map[string]int{
			"hello": 1,
			"world": 1,
			"and":   1,
			"love":  1,
			"thine": 1,
			"enemy": 1,
			"for":   1,
			"this":  1,
			"is":    1,
			"the":   2,
			"word":  1,
			"of":    1,
			"Lord":  1,
		}

		PerformMapAssertion(t, got, want)
	})
}

func PerformStringAssertion(t *testing.T, got, want string) {
	t.Helper() //this marks the function as a test helper function, so that when it fails, the line number reported will be in the caller of this function, rather than inside this function itself.
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func PerformMapAssertion(t *testing.T, got, want map[string]int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("got %v, want %v", got, want)
		return
	}
	for key, value := range want {
		if got[key] != value {
			t.Errorf("got %v, want %v", got, want)
			return
		}
	}
}
