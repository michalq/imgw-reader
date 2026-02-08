package climate

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func ParseSynopLine(l []string) (*Daily, error) {
	year, err := strconv.Atoi(l[2])
	if err != nil {
		return nil, fmt.Errorf("error parsing year: %w [%+v]", err, l)
	}
	month, err := strconv.Atoi(l[3])
	if err != nil {
		return nil, fmt.Errorf("error parsing month: %w", err)
	}
	day, err := strconv.Atoi(l[4])
	if err != nil {
		return nil, fmt.Errorf("error parsing day: %w", err)
	}
	maxTemp, err := parseFloat(l[5])
	if err != nil {
		return nil, fmt.Errorf("error parsing maxTemp: %w", err)
	}
	minTemp, err := parseFloat(l[7])
	if err != nil {
		return nil, fmt.Errorf("error parsing minTemp: %w", err)
	}
	avgTemp, err := parseFloat(l[9])
	if err != nil {
		return nil, fmt.Errorf("error parsing avgTemp: %w", err)
	}
	minGroundTemp, err := parseFloat(l[11])
	if err != nil {
		return nil, fmt.Errorf("error parsing minGroundTemp: %w", err)
	}
	precipitation, err := parseFloat(l[13])
	if err != nil {
		return nil, fmt.Errorf("error parsing precipitation: %w, (%+v)", err, l[13])
	}
	snowDepthCm, err := parseFloat(l[16])
	if err != nil {
		return nil, fmt.Errorf("error parsing snow depth: %w, (%+v)", err, l[16])
	}

	stationName, err := charmap.Windows1250.NewDecoder().
		Bytes([]byte(strings.TrimSpace(l[1])))
	if err != nil {
		return nil, fmt.Errorf("error parsing win1250 ecoded string: %w", err)
	}

	return &Daily{
		StationId:           l[0],
		StationName:         string(stationName),
		Year:                year,
		Month:               month,
		Day:                 day,
		MaxTemp:             float32(maxTemp),
		MaxTempStatus:       MeasurementStatus(l[6]),
		MinTemp:             float32(minTemp),
		MinTempStatus:       MeasurementStatus(l[8]),
		AvgTemp:             float32(avgTemp),
		AvgTempStatus:       MeasurementStatus(l[10]),
		MinGroundTemp:       float32(minGroundTemp),
		MinGroundTempStatus: MeasurementStatus(l[12]),
		Precipitation:       float32(precipitation),
		PrecipitationStatus: MeasurementStatus(l[14]),
		PrecipitationType:   PrecipitationType(l[15]),
		SnowDepthCm:         float32(snowDepthCm),
		SnowDepthStatus:     MeasurementStatus(l[17]),
	}, nil
}
