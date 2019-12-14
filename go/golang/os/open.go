package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// func os.Open(name string) (*os.File, error)
	f, err := os.Open("hoge.txt")
	if err != nil {
		return
	}
	// func bufio.NewReader(rd io.Reader) *bufio.Reader
	// os.File implements interface type io.Reader
	r := bufio.NewReader(f)
	for {
		// func (*bufio.Reader).ReadLine() (line []byte, isPrefix bool, err error)
		line, _, err := r.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}
