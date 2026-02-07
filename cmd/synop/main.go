package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/michalq/imgw/internal/persist"
	"github.com/michalq/imgw/internal/synop"

	"github.com/urfave/cli/v3"
)

const debug = true
const baseSynopData = "https://danepubliczne.imgw.pl/data/dane_pomiarowo_obserwacyjne/dane_meteorologiczne/dobowe/synop/"

func main() {
	ctx := context.Background()
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
				Name:  "scan",
				Usage: "Scans all IMGW data from 'dobowe/synop' and creates single CSV file.",
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
					outFolder := c.String("raw-dir")
					outFile := c.String("out")

					measurements, err := synop.SynopScanner(debug, outFolder)
					if err != nil {
						panic(err)
					}

					fmt.Printf("\nRead %d daily measurements!\n", len(measurements))
					fmt.Printf("\nStations\n")
					fmt.Printf("Id,Name,Oldest measurement,Newest measurement,Longest work\n")
					for _, s := range measurements.UniqueStations() {
						fmt.Printf("%s,%s,%s,%s,%d\n",
							s.Id,
							s.Name,
							s.OldestMeasurement.Format("2006-01-02"),
							s.NewestMeasurement.Format("2006-01-02"),
							int(s.LongestWork.Seconds()),
						)
					}

					persist.ToCsv(outFile, measurements)
					return nil
				},
			},
		},
	}
	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
