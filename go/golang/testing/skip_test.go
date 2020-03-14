package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	// <setup code>
	t.Run("A=1", func(t *testing.T) {
		t.Log("A=1")
	})
	t.Run("A=2", func(t *testing.T) {
		t.Log("A=2")
	})
	t.Run("B=1", func(t *testing.T) {
		t.Log("B=1")
	})
	// <tear-down code>
}
