package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Run(`
    	abc = 2 + 2;
    	console.log("The value of abc is " + abc); // 4
	`)
	value, err := vm.Get("abc")
	if err != nil {
		panic(err)
	}
	valueInt64, _ := value.ToInteger()
	fmt.Println(valueInt64)

	// vm.Run(`var a = {"b": "b", "c": undefined};`)
	vm.Run(`var a = {c: undefined};`)
	a, err := vm.Get("a")
	if err != nil {
		panic(err)
	}
	rawMap, _ := a.Export()
	// myMap := rawMap.(map[string]interface{})
	myMap := rawMap.(map[string]interface{})
	for k, v := range myMap {
		fmt.Println(k, v)
	}
	fmt.Printf("%#v\n", a.Object())
	fmt.Printf("%v\n", a.Object())
	fmt.Printf("%v\n", a.Object().Keys())
	v, _ := a.Object().Get("c")
	fmt.Printf("%v\n", v)
	fmt.Printf("%T\n", v)
	fmt.Printf("%t\n", v.IsUndefined())
	fmt.Printf("%d %d\n", len(a.Object().Keys()), len(myMap))
}
