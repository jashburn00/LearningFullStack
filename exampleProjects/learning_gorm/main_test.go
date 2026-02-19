package main

import (
	"os"
	"testing"
)

var originalStdout = os.Stdout

func TestMain(t *testing.T) {
	t.Run("Main should run without error", func(t *testing.T) {
		//setup
		r, w, _ := os.Pipe()
		defer r.Close()
		defer w.Close()
		os.Stdout = w

		//testing
		Main()

		//teardown
		os.Stdout = originalStdout
	})
}
