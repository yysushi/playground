package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}

// TODO: RunParallel, testing.PB
