package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func myFunc(long bool) {
	_ = fmt.Sprintf("hello")
	if long {
		time.Sleep(1 * time.Second)
	}
}

func BenchmarkHello(b *testing.B) {
	var i int
	log.Println(i, "started")
	for i = 0; i < b.N; i++ {
		myFunc(false)
	}
	log.Println(i, "finished")
}

func BenchmarkHelloLong(b *testing.B) {
	var i int
	log.Println(i, "started")
	for i = 0; i < b.N; i++ {
		myFunc(true)
	}
	log.Println(i, "finished")
}
