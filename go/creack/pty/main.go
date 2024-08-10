package main

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	c := exec.Command("echo", "bar\nfoobaz")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	var buf []byte = make([]byte, 2)
	func() {
		for {
			n, err := f.Read(buf)
			fmt.Printf("read: %d %v\n", n, err)
			if err == nil {
				fmt.Printf("out: %q\n", buf[:n])
				continue
			}
			if err == io.EOF || err.Error() == "read /dev/ptmx: input/output error" {
				return
			}
		}
	}()
}
