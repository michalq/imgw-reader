package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/michalq/imgw/internal/climate"
	"github.com/michalq/imgw/internal/crawler"
	"io"
	"os"
	"strconv"
)

const path = "/Users/michalkutrzeba/work/climate"

func main() {
	fmt.Println("Starting...")
	zipReader := crawler.NewZipReader(path + "/raw")
	climateData := crawler.NewClimate()
	climateParser := climate.NewDailyParser()

	measurements := make(climate.List, 0)
	for _, pckg := range climateData.Packages() {
		for _, file := range zipReader.Files(pckg.FileName) {
			if climate.IsClimateDailyFile(file.Name, pckg) {
				continue
			}
			rc, err := file.Open()
			if err != nil {
				panic(err)
			}
			data, err := io.ReadAll(rc)
			if err != nil {
				panic(err)
			}
			r := csv.NewReader(bytes.NewReader(data))
			acqs, err := climateParser.ParseFromReader(r)
			if err != nil {
				panic(err)
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

	f, err := os.Create(path + "/out/out.csv")
	if err != nil {
		panic(err)
	}
	wCsv := csv.NewWriter(f)
	for _, d := range measurements.ByStation("251200030") {
		if err := wCsv.Write([]string{
			d.StationId,
			d.Date(),
			strconv.FormatFloat(float64(d.AvgTemp), 'f', -1, 32),
		}); err != nil {
			panic(err)
		}
	}
	fmt.Println("\nMy job here is done, bye!")
}
