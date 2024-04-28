package climate

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
