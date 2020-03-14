package main

import (
	"fmt"
)

func Example() {
	fmt.Println(2)

	// Output:
	// 2
}

func M() string {
	return "2"
}

func ExampleM() {
	fmt.Println(M())

	// Output:
	// 2
}

type X struct{}

func ExampleX() {
	var _ X
	fmt.Println(2)

	// Output:
	// 2
}

func (x X) M() string {
	return "3"
}

func ExampleX_M() {
	var x X
	fmt.Println(x.M())

	// Output:
	// 3
}
