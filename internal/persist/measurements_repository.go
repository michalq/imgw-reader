package persist

import (
	"context"

	"github.com/michalq/imgw/internal/climate"
)

type MeasurementsRepository interface {
	Setup(ctx context.Context) error
	AddMeasurements(ctx context.Context, measurements []climate.Daily) error
	AddStations(ctx context.Context, stations []climate.Station) error
}
