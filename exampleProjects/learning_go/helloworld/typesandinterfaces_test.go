package main

import (
	"testing"
)

func TestDopeList(t *testing.T) {

	t.Run("should support adding items", func(t *testing.T) {
		stringList := DopeList[string]{}
		stringList.Add("Hello")
		stringList.Add("World")

		if len(stringList.items) != 2 {
			t.Errorf("want 2 items, got %d", len(stringList.items))
		}
	})

	t.Run("should support different types", func(t *testing.T) {
		intList := DopeList[int]{}
		intList.Add(1)
		intList.Add(2)

		if len(intList.items) != 2 {
			t.Errorf("want 2 items, got %d", len(intList.items))
		}
	})
}
