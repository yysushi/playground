package hcl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

func TestSimple(t *testing.T) {
	var config Config
	err := hclsimple.DecodeFile("testdata/read.hcl", nil, &config)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("Configuration is %#v", config)

}
