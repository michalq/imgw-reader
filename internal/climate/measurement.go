package climate

type (
	MeasurementStatus string
	PrecipitationType string
)

const (
	// NoMeasurement Status "8" brak pomiaru
	NoMeasurement MeasurementStatus = "8"
	// NoSmth Status "9" brak zjawiska
	NoSmth MeasurementStatus = "9"

	PrecipitationTypeS PrecipitationType = "S"
	PrecipitationTypeW PrecipitationType = "W"
)
