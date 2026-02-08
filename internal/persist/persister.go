package persist

import (
	"context"

	"github.com/michalq/imgw/internal/climate"
)

type Persister interface {
	Setup(ctx context.Context) error
	Persist(ctx context.Context, measurements climate.List) error
	PersistStations(ctx context.Context, stations []climate.Station) error
}
