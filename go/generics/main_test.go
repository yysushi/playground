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

type Validator interface {
	Validate(string) bool
}

type Provider[T Validator] struct {
	// NOTE: embedded field type cannot be a (pointer to a) type parameter (typecheck)
	// T
	validator T
}

// check "known limitations" in https://tip.golang.org/doc/go1.18#generics
// and this https://github.com/golang/go/issues/49030
func (p Provider[T]) Validate(input string) bool {
	return p.validator.Validate(input)
}

func (p Provider[T]) Create(input string) bool {
	if !p.Validate(input) {
		return false
	}
	// some processings...
	return true
}

type NetworkResource struct{}

func (r NetworkResource) Validate(input string) bool {
	return len(input) < 5
}

func TestA(t *testing.T) {
	p := &Provider[NetworkResource]{
		validator: NetworkResource{},
	}
	if p.Create("12345") {
		t.Fail()
	}
}
