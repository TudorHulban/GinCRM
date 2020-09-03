package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TudorHulban/GinCRM/pkg/httpinterface"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"github.com/TudorHulban/GinCRM/pkg/vers"
	tlog "github.com/TudorHulban/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		fmt.Println(vers.Version) // fmt used instead of log for nicer output.
		os.Exit(0)
	}

	ctx, cancel := ostop.NewOSCancellableCtx()
	defer cancel()

	// creating an error group to keep dependencies in sync, only Gin dependency now though.
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		cfg, errConfig := httpinterface.CreateConfig("0.0.0.0:8080", "0.1", tlog.DEBUG, 3)
		if errConfig != nil {
			return errors.WithMessage(errConfig, "errors creating HTTP Server configuration")
		}
		gin, errCo := httpinterface.NewGinServer(cfg)
		if errCo != nil {
			return errors.WithMessage(errCo, "could not create HTTP Server instance")
		}
		return gin.Run(ctx)
	})

	if errWait := g.Wait(); errWait != nil {
		log.Println("Error group runtime error: ", errWait)
		os.Exit(ostop.RUNTIME)
	}
}
