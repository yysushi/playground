package main

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Type string
	Name string
}

func main() {
	var result Person
	input := map[string]string{
		"type": "person",
		"name": "koke",
	}
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}
