package main_test

import (
	"fmt"
	"testing"
)

func Sum[T int | float64](t []T) T {
	var sum T
	for _, tt := range t {
		sum += tt
	}
	return sum
}

var sumInt = Sum[int]

func TestA(t *testing.T) {
	if Sum[int]([]int{1, 23}) != 24 {
		t.Fail()
	}
	if sumInt([]int{1, 2, 3}) != 6 {
		t.Fail()
	}
}

type Animal interface {
	Bark() string
}

func SayHello[T Animal](t T) string {
	return fmt.Sprintf("%s Hello!", t.Bark())
}

type Dog struct{}

func (Dog) Bark() string {
	return "Woof!"
}

// create
// update
// show
// list
// delete

// create network <- create & network impl
// create network <- create & network impl

func TestB(t *testing.T) {
	dog := Dog{}
	if SayHello[Dog](dog) != "Woof! Hello!" {
		t.Fail()
	}
}

type Stringer interface {
	String() string
}

type A struct {
}
func (a A) string() string{
}
A.String = A.string
