package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("start")
	// call flag.Parse() here if TestMain uses flags
	defer log.Println("end")
	// os.Exit(m.Run())
	m.Run()
}
