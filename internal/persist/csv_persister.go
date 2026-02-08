package persist

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/michalq/imgw/internal/climate"
)

type CsvPersister struct {
	path string
}

func NewCsvPersister(path string) *CsvPersister {
	return &CsvPersister{path}
}

func (c *CsvPersister) Setup(ctx context.Context) error {
	return nil
}

func (c *CsvPersister) Persist(
	_ context.Context,
	measurements climate.List,
) error {
	f, err := os.Create(c.path)
	if err != nil {
		panic(err)
	}
	wCsv := csv.NewWriter(f)
	for _, d := range measurements {
		if err := wCsv.Write([]string{
			d.StationId,
			d.Date(),
			strconv.Itoa(d.Year),
			strconv.Itoa(d.Month),
			strconv.Itoa(d.Day),
			strconv.FormatFloat(float64(d.AvgTemp), 'f', -1, 32),
			strconv.FormatFloat(float64(d.MaxTemp), 'f', -1, 32),
			strconv.FormatFloat(float64(d.MinTemp), 'f', -1, 32),
		}); err != nil {
			panic(err)
		}
	}
	wCsv.Flush()
	return nil
}

func (c *CsvPersister) PersistStations(_ context.Context, _ []climate.Station) error {
	fmt.Println("CSV stations persister not implemented.")
	return nil
}
