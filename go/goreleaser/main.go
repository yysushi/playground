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

	cfg, err := config.LoadReader(fd)
	if err != nil {
		panic(err)
	}

	cfg.Release.GitHub.Owner = "goreleaser"
	cfg.Release.GitHub.Name = "goreleaser"

	var ctx = rcontext.New(cfg)

	for _, defaulter := range defaults.Defaulters {
		if err := defaulter.Default(ctx); err != nil {
			fmt.Printf("%#v\n", defaulter)
			panic(err)
		}
	}
	fmt.Printf("%#v\n", ctx.Config.Archives[0].NameTemplate)
}
