package io

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadFile1(t *testing.T) {
	// func os.Open(name string) (*os.File, error)
	f, err := os.Open("./testdata/read_hoge.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	// func ioutil.ReadAll(r io.Reader) ([]byte, error)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error(err)
	}
	expected := "hello world\nkoketani\n\n"
	if string(b) != expected {
		t.Errorf("expected %s, but %s", expected, b)
	}
}

func TestReadFile2(t *testing.T) {
	// func os.Open(name string) (*os.File, error)
	f, err := os.Open("./testdata/read_hoge.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// func bufio.NewReader(rd io.Reader) *bufio.Reader
	r := bufio.NewReader(f)
	for {
		// func (*bufio.Reader).ReadLine() (line []byte, isPrefix bool, err error)
		a, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error(err)
		}
		t.Log(string(a))
	}
}

func TestReadFile3(t *testing.T) {
	// func os.Open(name string) (*os.File, error)
	f, err := os.Open("./testdata/read_hoge.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	// func bufio.NewReader(rd io.Reader) *bufio.Reader
	r := bufio.NewReader(f)
	for {
		// func (*bufio.Reader).ReadString(delim byte) (string, error)
		a, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error(err)
		}
		t.Log(a)
	}
}

func TestReadFile4(t *testing.T) {
	// func ioutil.ReadFile(filename string) ([]byte, error)
	b, err := ioutil.ReadFile("./testdata/read_hoge.txt")
	if err != nil {
		t.Error(err)
	}
	expected := "hello world\nkoketani\n\n"
	if string(b) != expected {
		t.Errorf("expected %s, but %s", expected, b)
	}
}
