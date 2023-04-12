package main

import (
	"embed"
	"fmt"
)

//go:embed 1.txt
var hoge1 string

//go:embed 2.txt
var hoge2 []byte

//go:embed 3/3.txt
var hoge3 embed.FS

func main() {
	fmt.Println(hoge1)
	fmt.Println(hoge2)
	d, err := hoge3.ReadFile("3/3.txt")
	fmt.Println(err, d)
}
