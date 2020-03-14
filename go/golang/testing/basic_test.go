package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	if 1+1 != 2 {
		t.Fatalf("expected 2, but %v", 1+1)
	}
}
