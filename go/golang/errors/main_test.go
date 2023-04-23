package main_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errNew = errors.New("errors.New")
var errFmt = fmt.Errorf("fmt.Errorf")
var errFmtWrapped = fmt.Errorf("outer wrapping error: %w", errFmt)
var errJoined = errors.Join(errNew, errors.New("later joined error"))

func TestErrorMessage(t *testing.T) {
	// output
	assert.Errorf(t, errNew, "errors.New")
	assert.Errorf(t, errFmt, "fmt.Errorf")
	assert.ErrorContains(t, errFmtWrapped, "outer wrapping error: fmt.Errorf")
	assert.ErrorContains(t, errJoined, "errors.New\nlater joined error")

	// errors.Unwrap
	assert.Equal(t, errors.Unwrap(errFmtWrapped), errFmt)
	// assert.Equal(t, errNew, errors.Unwrap(errJoined))

	// error.Is
	// two ways of comparing errors
	// 1. ==
	// 2. errors.Is
	// the later "errors.Is" is preferable in that
	// it can checked in internal error
	assert.NotEqual(t, errFmtWrapped, errFmt)
	assert.ErrorIs(t, errFmtWrapped, errFmt)
	assert.ErrorIs(t, fmt.Errorf("another outer error: %w", errFmtWrapped), errFmt)
	assert.ErrorIs(t, errFmt, errFmt)
	assert.ErrorIs(t, errJoined, errNew)

	// errors.As
	var perr *os.PathError
	assert.ErrorAs(t, fmt.Errorf("outer error: %w", &os.PathError{Path: "somewhere"}), &perr)
}

func Example_type() {
	fmt.Printf("%T, %T, %T, %T, %T\n", errNew, errFmt, errFmtWrapped, errors.Unwrap(errFmtWrapped), errJoined)
	// Output: *errors.errorString, *errors.errorString, *fmt.wrapError, *errors.errorString, *errors.joinError
}
