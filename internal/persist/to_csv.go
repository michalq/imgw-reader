package persist

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/michalq/imgw/internal/climate"
)

func ToCsv(path string, measurements climate.List) {
	f, err := os.Create(path + "/out/out.csv")
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
}
