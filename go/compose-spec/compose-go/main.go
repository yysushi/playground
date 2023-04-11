package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

func main() {
	// 1. load original file
	content, err := os.ReadFile("./compose.yml")
	if err != nil {
		panic(err)
	}
	var project *types.Project
	project, err = loader.Load(types.ConfigDetails{
		ConfigFiles: []types.ConfigFile{{Content: content}},
		Environment: map[string]string{},
	})
	if err != nil {
		panic(err)
	}
	// 2. manipulate the project
	//    - filter only interesting service
	//    - update the service network on the host
	var targetImageName string = "redis"
	var filteredSvcs types.Services
	for _, svc := range project.AllServices() {
		if !strings.Contains(svc.Image, targetImageName) {
			continue
		}
		// svc.NetworkMode = "host"
		svc.Networks = map[string]*types.ServiceNetworkConfig{"host": nil}
		filteredSvcs = append(filteredSvcs, svc)
	}
	project.Services = filteredSvcs
	project.WithoutUnnecessaryResources()

	// 3. dump to stdout
	var b []byte
	b, err = project.MarshalYAML()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}
