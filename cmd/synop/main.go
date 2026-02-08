package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/michalq/imgw/internal/adapters"
	"github.com/michalq/imgw/internal/config"
	"github.com/michalq/imgw/internal/persist"
	"github.com/michalq/imgw/internal/synop"

	"github.com/urfave/cli/v3"
)

const debug = true
const baseSynopData = "https://danepubliczne.imgw.pl/data/dane_pomiarowo_obserwacyjne/dane_meteorologiczne/dobowe/synop/"
const configPath = "common_config.json"

func main() {
	ctx := context.Background()
	cfg, err := config.Read(configPath)
	if err != nil {
		panic(err)
	}

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "download",
				Usage: "Downloads all IMGW data from 'dobowe/synop'.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "raw-dir",
						Value:    "",
						Required: true,
						Usage:    "Directory where to save raw data",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					outFolder := c.String("raw-dir")
					if err := synop.SynopDataDownloader(debug, baseSynopData, outFolder); err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:  "import",
				Usage: "Scans all IMGW data from 'dobowe/synop' and saves to csv or database.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "raw-dir",
						Value:    "",
						Required: true,
						Usage:    "Directory with raw data",
					},
					&cli.StringFlag{
						Name:     "out",
						Value:    "",
						Required: true,
						Usage:    "Output file",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					// Input
					outFolder := c.String("raw-dir")
					outFile := c.String("out")
					// Dependencies
					sqliteDb, err := adapters.ProviderSqlite(cfg.Db.Url)
					if err != nil {
						panic(err)
					}
					measurementsRepo := adapters.NewSqliteMeasurementsRepository(sqliteDb)
					persisters := []persist.Persister{
						persist.NewCsvPersister(outFile),
						persist.NewDatabasePersister(measurementsRepo),
					}
					// Import & persist
					measurements, err := synop.SynopScanner(debug, outFolder)
					if err != nil {
						log.Fatalf("failed to scan measurements: %v", err)
					}
					fmt.Printf("\nRead %d daily measurements!\n", len(measurements))
					stations := measurements.UniqueStations()
					for _, p := range persisters {
						if err := p.Setup(ctx); err != nil {
							log.Fatalf("cannot setup persister: %+v", err)
						}
						if err := p.Persist(ctx, measurements); err != nil {
							log.Fatalf("cannot persist measurements %+v", err)
						}
						if err := p.PersistStations(ctx, stations); err != nil {
							log.Fatalf("cannot persist measurements %+v", err)
						}
					}
					return nil
				},
			},
		},
	}
	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
