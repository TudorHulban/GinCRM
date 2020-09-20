package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TudorHulban/GinCRM/cmd/settings"

	"github.com/TudorHulban/GinCRM/pkg/httpinterface"
	"github.com/TudorHulban/GinCRM/pkg/ostop"
	"github.com/TudorHulban/GinCRM/pkg/persistence/cgorm"
	tlog "github.com/TudorHulban/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		fmt.Println(settings.Version) // fmt used instead of log for nicer output.
		os.Exit(0)
	}

	ctx, cancel := ostop.NewOSCancellableCtx()
	defer cancel()

	// clean up
	if _, errExists := os.Stat("/path/to/whatever"); !os.IsNotExist(errExists) {
		if errCleanSQLite := os.Remove(settings.SQLiteFilePath); errCleanSQLite != nil {
			log.Println("Error removing previous SQLite database file, error: ", errCleanSQLite)
			os.Exit(ostop.SQLiteCleanUp)
		}
	}

	// creating an error group to keep dependencies in sync, only Gin dependency now though.
	g, ctx := errgroup.WithContext(ctx)

	// creating RDBMS schema
	g.Go(func() error {
		return cgorm.MigrateDBSchema()
	})

	// creating HTTP layer
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
