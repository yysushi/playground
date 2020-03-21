package main

import (
	"flag"
	"fmt"
)

func main() {
	a := flag.String("config", "config.sample.json", "")
	// var a string
	// flag.String(&a, "config", "config.sample.json", "")
	flag.Parse()
	fmt.Println(a)
}
