package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"

	_ "github.com/lib/pq"

	"github.com/conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/backend"
	"github.com/fieldkit/cloud/server/backend/handlers"
)

type Options struct {
	StationID int
}

type Config struct {
	PostgresURL string `split_words:"true" default:"postgres://localhost/fieldkit?sslmode=disable" required:"true"`
}

func main() {
	options := &Options{}

	flag.IntVar(&options.StationID, "station-id", 0, "station id")

	flag.Parse()

	if options.StationID == 0 {
		flag.Usage()
		os.Exit(2)
	}

	config := &Config{}
	if err := envconfig.Process("FIELDKIT", config); err != nil {
		panic(err)
	}

	db, err := sqlxcache.Open("postgres", config.PostgresURL)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	err = db.WithNewTransaction(ctx, func(txCtx context.Context) error {
		rw := backend.NewRecordWalker(db)

		started := time.Now()

		log.Printf("processing\n")

		walkParams := &backend.WalkParameters{
			StationID: int32(options.StationID),
			Start:     time.Time{},
			End:       time.Now(),
			PageSize:  1000,
		}

		visitor := handlers.NewAggregatingHandler(db)
		if err := rw.WalkStation(txCtx, visitor, walkParams); err != nil {
			panic(err)
		}

		info, err := rw.Info(txCtx)
		if err != nil {
			panic(err)
		}

		finished := time.Now()

		log.Printf("done %v data (%v meta) %v\n", info.DataRecords, info.MetaRecords, finished.Sub(started))

		return nil
	})
	if err != nil {
		panic(err)
	}
}