package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("should add two positive numbers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5
		performIntegerAssertion(t, got, want)
	})
	t.Run("should add a positive and a negative number", func(t *testing.T) {
		got := Add(5, -2)
		want := 3
		performIntegerAssertion(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}

func performIntegerAssertion(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
