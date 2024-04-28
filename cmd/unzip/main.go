package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/michalq/imgw/internal/climate"
	"github.com/michalq/imgw/internal/crawler"
	"io"
)

const path = "/Users/michalkutrzeba/work/climate/raw"

func main() {
	fmt.Println("Starting...")
	zipReader := crawler.NewZipReader(path)
	climateData := crawler.NewClimate()
	climateParser := climate.NewDailyParser()
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
			fmt.Printf("Read %d daily measuremenets from package %s\n", len(acqs), pckg.FileName)
		}
	}
	fmt.Println("My job here is done, bye!")
}
