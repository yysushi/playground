package main

import "fmt"

type myEntry struct{}

func (e *myEntry) String() string {
	return "entry"
}

type myType interface {
	String() string
}

type Server[T1 myType, T2 []T1] struct {
	t T2
}

func (s *Server[T1, T2]) Hoge(T1) {
}

type myTypedServer = Server[*myEntry, []*myEntry]

func main() {
	s := &myTypedServer{
		t: []*myEntry{&myEntry{}},
	}
	fmt.Printf("%#v", s)
	fmt.Println()
}
