package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v61/github"
	"github.com/goreleaser/goreleaser/pkg/config"
	rcontext "github.com/goreleaser/goreleaser/pkg/context"
	"github.com/goreleaser/goreleaser/pkg/defaults"
)

func main() {
	client := github.NewClient(nil)
	fd, _, err := client.Repositories.DownloadContents(context.Background(), "goreleaser", "goreleaser", ".goreleaser.yaml", &github.RepositoryContentGetOptions{Ref: "v1.25.1"})
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// buf := new(strings.Builder)
	// _, err = io.Copy(buf, fd)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(buf)

	cfg, err := config.LoadReader(fd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", cfg.Archives)

	var ctx = rcontext.New(cfg)

	for _, defaulter := range defaults.Defaulters {
		if err := defaulter.Default(ctx); err != nil {
			fmt.Printf("%#v\n", defaulter)
			panic(err)
		}
	}
}
