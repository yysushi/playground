package main

import "fmt"

type Hoge struct {
	hoge int
}

type Fuga struct {
	Hoge
	fuga string
}

type Fuga2 struct {
	*Hoge
	fuga string
}

func main() {
	h := Hoge{1}
	fmt.Printf("%p\n", &h.hoge)
	f := Fuga{h, "2"}
	fmt.Printf("%p\n", &f.hoge)
	f2 := Fuga2{&h, "2"}
	fmt.Printf("%p\n", &f2.hoge)
}
