package persist

import (
	"context"

	"github.com/michalq/imgw/internal/climate"
)

type DatabasePersister struct {
	repo MeasurementsRepository
}

func NewDatabasePersister(repo MeasurementsRepository) *DatabasePersister {
	return &DatabasePersister{repo}
}

func (d *DatabasePersister) Setup(ctx context.Context) error {
	err := d.repo.Setup(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabasePersister) Persist(ctx context.Context, measurements climate.List) error {
	return d.repo.AddMeasurements(ctx, measurements)
}

func (d *DatabasePersister) PersistStations(ctx context.Context, stations []climate.Station) error {
	return d.repo.AddStations(ctx, stations)
}
