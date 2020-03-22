package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// New
	err1 := errors.New("db error")
	err2 := fmt.Errorf("db error")
	fmt.Printf("err1:`%s`, err2:`%s`\n", err1, err2)
	// Wrap
	err3 := fmt.Errorf("my app error: %w", err2)
	fmt.Printf("err3:`%s`\n", err3)
	// Unwrap
	err4 := errors.Unwrap(err3)
	fmt.Printf("err4:`%s`\n", err4)
	// func errors.Is(err error, target error) bool
	// is preferable to the below, because the former will succeed if err wraps os.ErrExist.
	// if err == os.ErrExist
	fmt.Printf("errors.Is(err3, err4):`%t`\n", errors.Is(err3, err4))
	fmt.Printf("err3==err4:`%t`\n", err3 == err4)
	fmt.Printf("errors.Is(err3, err2):`%t`\n", errors.Is(err3, err2))
	fmt.Printf("errors.Is(err2, err3):`%t`\n", errors.Is(err2, err3))
	fmt.Printf("errors.Is(err2, err4):`%t`\n", errors.Is(err2, err4))
	fmt.Printf("err2==err4:`%t`\n", err2 == err4)
	// func errors.As(err error, target interface{}) bool
	// is preferable to below, because the former will succeed if err wraps an *os.PathError.
	// if perr, ok := err.(*os.PathError); ok {
	// 	fmt.Println(perr.Path)
	// }
	var perr *os.PathError
	if ok := errors.As(fmt.Errorf("my error: %w", &os.PathError{Path: "somewhere"}), &perr); ok {
		fmt.Printf("%s\n", perr.Path)
	}
}
