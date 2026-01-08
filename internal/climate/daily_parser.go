package climate

import (
	"encoding/csv"
	"io"
	"strconv"
)

type DailyParser struct {
	parseLine func(l []string) (*Daily, error)
}

func NewDailyParser(parseLine func(l []string) (*Daily, error)) DailyParser {
	return DailyParser{parseLine}
}

func (r DailyParser) ParseFromReader(reader *csv.Reader) ([]Daily, error) {
	acqs := make([]Daily, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		l, err := r.parseLine(record)
		if err != nil {
			return nil, err
		}
		acqs = append(acqs, *l)
	}
	return acqs, nil
}

func (r DailyParser) ParseAll(all [][]string) ([]Daily, error) {
	acqs := make([]Daily, 0, len(all))
	for _, l := range all {
		acq, err := r.parseLine(l)
		if err != nil {
			return nil, err
		}
		acqs = append(acqs, *acq)
	}
	return acqs, nil
}

func parseFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 32)

}
