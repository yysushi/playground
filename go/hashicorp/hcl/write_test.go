package hcl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestSimple2(t *testing.T) {
	var config Config
	err := hclsimple.DecodeFile("testdata/read.hcl", nil, &config)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&config, f.Body())
	fmt.Printf("%s", f.Bytes())

}
