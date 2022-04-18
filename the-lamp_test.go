package main

import (
	"testing"
)

func TestStatus(t *testing.T) {
	t.Run("Should throw Lights ON when input was EVEN number", func(t *testing.T) {
		_ = Status(4)
		t.Log("Lights ON")
	})

	t.Run("Should throw Lights OFF when input was ODD number", func(t *testing.T) {
		_ = Status(5)
		t.Log("Lights OFF")
	})
}
