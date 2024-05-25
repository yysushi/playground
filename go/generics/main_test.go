package main_test

import (
	"testing"
)

// create
// update
// show
// list
// delete

// create network <- create & network impl
// create network <- create & network impl

type Logic struct {
	validate func(string) bool
}

type Provider struct {
	Logic
}

func (p Provider) Create(input string) bool {
	if !p.validate(input) {
		return false
	}
	// some processings...
	return true
}

func networkValidate(input string) bool {
	return len(input) < 5
}

func TestA(t *testing.T) {
	p := &Provider{
		Logic{
			validate: networkValidate,
		},
	}
	if p.Create("12345") {
		t.Fail()
	}
}
