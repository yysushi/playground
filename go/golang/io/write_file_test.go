package io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteFile1(t *testing.T) {
	// func os.Open(name string) (*os.File, error)
	f, err := os.Open("./testdata/write_hoge.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// func bufio.NewWriter(w io.Writer) *bufio.Writer
	w := bufio.NewWriter(f)
	_, err = w.WriteString("aaa")
	if err != nil {
		t.Error(err)
	}
	// func (*bufio.Writer).WriteString(s string) (int, error)
	_, err = w.WriteString("aaa")
	if err != nil {
		t.Error(err)
	}
	// func io.WriteString(w io.Writer, s string) (n int, err error)
	io.WriteString(w, "bbb")
}

func TestWriteFile2(t *testing.T) {
	// func ioutil.WriteFile(filename string, data []byte, perm os.FileMode) error
	err := ioutil.WriteFile("./testdata/write_hoge.txt", []byte("aaa"), 0600)
	if err != nil {
		t.Error(err)
	}
}

func TestWriteFile3(t *testing.T) {
	f, err := os.Open("./testdata/write_hoge.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	// func fmt.Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	fmt.Fprintf(w, "ccc")
}
