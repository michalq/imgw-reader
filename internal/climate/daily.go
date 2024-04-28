package climate

import (
	"fmt"
	"time"
)

type List []Daily

func (l List) ByStation(stationId string) List {
	list := make(List, 0)
	for _, d := range l {
		if d.StationId == stationId {
			list = append(list, d)
		}
	}
	return list
}

func (l List) UniqueStations() []Station {
	stationsMap := make(map[string]*Station)
	for _, d := range l {
		if station, ok := stationsMap[d.StationId]; ok {
			if d.Time().Before(station.OldestMeasurement) {
				station.OldestMeasurement = d.Time()
			}
			if d.Time().After(station.NewestMeasurement) {
				station.NewestMeasurement = d.Time()
			}
			station.LongestWork = station.NewestMeasurement.Sub(station.OldestMeasurement)
		} else {
			station := Station{
				Id:                d.StationId,
				Name:              d.StationName,
				OldestMeasurement: d.Time(),
				NewestMeasurement: d.Time(),
				LongestWork:       0 * time.Second,
			}
			stationsMap[d.StationId] = &station
		}
	}

	stations := make([]Station, 0)
	for _, s := range stationsMap {
		stations = append(stations, *s)
	}
	return stations
}

type Station struct {
	Id                string
	Name              string
	OldestMeasurement time.Time
	NewestMeasurement time.Time
	LongestWork       time.Duration
}
type Daily struct {
	// StationId Kod stacji
	StationId string
	// StationName Nazwa stacji
	StationName string
	// Year Rok
	Year int
	// Month Miesiąc
	Month int
	// Day Dzień
	Day int
	// MaxTemp Maksymalna temperatura dobowa [°C]
	MaxTemp float32
	// MaxTempStatus Status pomiaru TMAX                     1
	MaxTempStatus MeasurementStatus
	// MinTemp Minimalna temperatura dobowa [°C]
	MinTemp float32
	// MinTempStatus Status pomiaru TMIN
	MinTempStatus MeasurementStatus
	// AvgTemp Średnia temperatura dobowa [°C]
	AvgTemp float32
	// AvgTempStatus Status pomiaru STD
	AvgTempStatus MeasurementStatus
	// MinGroundTemp Temperatura minimalna przy gruncie [°C]
	MinGroundTemp float32
	// MinGroundTempStatus Status pomiaru TMNG
	MinGroundTempStatus MeasurementStatus
	// Precipitation Suma dobowa opadów [mm]
	Precipitation float32
	// PrecipitationStatus Status pomiaru SMDB                     1
	PrecipitationStatus MeasurementStatus
	// PrecipitationType Rodzaj opadu  [S/W/ ]
	PrecipitationType PrecipitationType
	// SnowHeight Wysokość pokrywy śnieżnej [cm]
	SnowHeight float32
	// SnowHeightStatus Status pomiaru PKSN                     1
	SnowHeightStatus MeasurementStatus
}

// Date formatted in Y-M-D
func (d Daily) Date() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d Daily) Time() time.Time {
	t, _ := time.Parse("2006-01-02", d.Date())
	return t
}
