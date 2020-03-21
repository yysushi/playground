package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// func os.Create(name string) (*os.File, error)
	wf, err := os.Create("hoge.txt")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(wf)
	// func fmt.Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	fmt.Fprintln(w, "hello\nworld")
	w.WriteString("hello world again\n")
	wf.Close()
	// func bufio.NewReader(rd io.Reader) *bufio.Reader
	// os.File implements interface type io.Reader
	// func os.Open(name string) (*os.File, error)
	rf, err := os.Open("hoge.txt")
	if err != nil {
		panic(err)
	}
	// func bufio.NewReader(rd io.Reader) *bufio.Reader
	// os.File implements interface type io.Reader
	r := bufio.NewReader(rf)
	for {
		// func (*bufio.Reader).ReadLine() (line []byte, isPrefix bool, err error)
		line, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			fmt.Println(string(line))
			break
		}
		fmt.Println(string(line))
	}
}
