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
var errFmtWrapped = fmt.Errorf("outer error: %w", errFmt)

func TestErrorMessage(t *testing.T) {
	assert.Errorf(t, errNew, "errors.New")
	assert.Errorf(t, errFmt, "fmt.Errorf")
	assert.Errorf(t, errFmtWrapped, "outer error: fmt.Errorf")
	assert.Equal(t, errFmt, errors.Unwrap(errFmtWrapped))
	// errors.Is
	// if wrapped or equal, errors.Is passed without any error
	assert.NotEqual(t, errFmt, errFmtWrapped)
	assert.ErrorIs(t, errFmtWrapped, errFmt)
	assert.ErrorIs(t, fmt.Errorf("another outer error: %w", errFmtWrapped), errFmt)
	assert.ErrorIs(t, errFmt, errFmt)
	// errors.As
	var perr *os.PathError
	assert.ErrorAs(t, fmt.Errorf("outer error: %w", &os.PathError{Path: "somewhere"}), &perr)
}

func Example_type() {
	fmt.Printf("%T, %T, %T, %T\n", errNew, errFmt, errFmtWrapped, errors.Unwrap(errFmtWrapped))
	// Output: *errors.errorString, *errors.errorString, *fmt.wrapError, *errors.errorString
}
