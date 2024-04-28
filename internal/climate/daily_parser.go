package climate

import (
	"encoding/csv"
	"io"
	"strconv"
)

type DailyParser struct{}

func NewDailyParser() DailyParser {
	return DailyParser{}
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
		l, err := r.ParseLine(record)
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
		acq, err := r.ParseLine(l)
		if err != nil {
			return nil, err
		}
		acqs = append(acqs, *acq)
	}
	return acqs, nil
}

func (r DailyParser) ParseLine(l []string) (*Daily, error) {
	year, err := strconv.Atoi(l[2])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(l[3])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(l[4])
	if err != nil {
		return nil, err
	}
	maxTemp, err := strconv.ParseFloat(l[5], 32)
	if err != nil {
		return nil, err
	}
	minTemp, err := strconv.ParseFloat(l[7], 32)
	if err != nil {
		return nil, err
	}
	avgTemp, err := strconv.ParseFloat(l[9], 32)
	if err != nil {
		return nil, err
	}
	minGroungTemp, err := strconv.ParseFloat(l[11], 32)
	if err != nil {
		return nil, err
	}
	precipitation, err := strconv.ParseFloat(l[13], 32)
	if err != nil {
		return nil, err
	}
	snowHeight, err := strconv.ParseFloat(l[16], 32)
	if err != nil {
		return nil, err
	}
	return &Daily{
		StationId:           l[0],
		StationName:         l[1],
		Year:                year,
		Month:               month,
		Day:                 day,
		MaxTemp:             float32(maxTemp),
		MaxTempStatus:       MeasurementStatus(l[6]),
		MinTemp:             float32(minTemp),
		MinTempStatus:       MeasurementStatus(l[8]),
		AvgTemp:             float32(avgTemp),
		AvgTempStatus:       MeasurementStatus(l[10]),
		MinGroundTemp:       float32(minGroungTemp),
		MinGroundTempStatus: MeasurementStatus(l[12]),
		Precipitation:       float32(precipitation),
		PrecipitationStatus: MeasurementStatus(l[14]),
		PrecipitationType:   PrecipitationType(l[15]),
		SnowHeight:          float32(snowHeight),
		SnowHeightStatus:    MeasurementStatus(l[17]),
	}, nil
}
