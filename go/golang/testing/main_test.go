package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("setup")
	// call flag.Parse() here if TestMain uses flags
	defer log.Println("teardown")
	// os.Exit(m.Run())
	m.Run()
}

func Test1(t *testing.T) {
	log.Println("1")
}

func Test2(t *testing.T) {
	log.Println("2")
}
