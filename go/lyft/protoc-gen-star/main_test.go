package main_test

import (
	"testing"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

func TestA(t *testing.T) {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(&myPGSModule{}).
		RegisterPostProcessor(&myPostProcessor{}).
		Render()
}
