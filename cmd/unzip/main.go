package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/michalq/imgw/internal/climate"
	"github.com/michalq/imgw/internal/crawler"
	"github.com/michalq/imgw/internal/persist"
)

const path = "/Users/michalkutrzeba/work/climate"

func main() {
	fmt.Println("Starting...")
	zipReader := crawler.NewZipReader(path + "/raw")
	climateData := crawler.NewClimate()
	climateParser := climate.NewDailyParser(climate.ParseClimateLine)

	errs := make([]error, 0)

	measurements := make(climate.List, 0)
	for _, pckg := range climateData.Packages() {
		for _, file := range zipReader.Files(pckg.FileName) {
			if !climate.IsClimateDailyFile(file.Name, pckg) {
				continue
			}
			fmt.Println("Reading ", file.Name)
			rc, err := file.Open()
			if err != nil {
				errs = append(errs, fmt.Errorf("could not open file %s (%w)", file.Name, err))
				continue
			}
			data, err := io.ReadAll(rc)
			if err != nil {
				errs = append(errs, fmt.Errorf("problem with reading file %s (%s)", file.Name, err))
				continue
			}
			r := csv.NewReader(bytes.NewReader(data))
			acqs, err := climateParser.ParseFromReader(r)
			if err != nil {
				errs = append(errs, fmt.Errorf("problem with parsing file %s (%s)", file.Name, err))
				continue
			}
			measurements = append(measurements, acqs...)
			fmt.Printf("Read %d daily measuremenets from package %s\n", len(acqs), pckg.FileName)
		}
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

	persist.ToCsv(path, measurements)

	if len(errs) > 0 {
		fmt.Printf("\nErrors encountered:\n")
		for _, e := range errs {
			fmt.Println(e)
		}
	}

	fmt.Println("\nMy job here is done, bye!")
}
