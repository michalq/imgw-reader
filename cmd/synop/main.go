package main

import (
	"fmt"
	"log"

	"github.com/michalq/imgw/internal/persist"
	"github.com/michalq/imgw/internal/synop"
)

const debug = true

const outFolder = "./raw/synop/"
const baseSynopData = "https://danepubliczne.imgw.pl/data/dane_pomiarowo_obserwacyjne/dane_meteorologiczne/dobowe/synop/"

const stepDownloadAll = false

func main() {
	if stepDownloadAll {
		if err := synop.SynopDataDownloader(debug, baseSynopData, outFolder); err != nil {
			log.Fatal(err)
		}
	}

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

	persist.ToCsv("./", measurements)

	// 4. scan

}
