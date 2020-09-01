package main

import (
	"log"

	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"github.com/docker/docker/daemon/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	// creating an error group to keep dependencies in sync, only Gin dependency now though.g.Go(func() error {
			return httpinterface.NewHTTPServer(address).Run(ctx, rtbrick.RTBrick{Cfg: config.New()})
		})

	ctx, cancel := ostop.NewOSCancellableCtx()
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return (address).Run(ctx, rtbrick.RTBrick{Cfg: config.New()})
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err, "runtime error")
	}
}
