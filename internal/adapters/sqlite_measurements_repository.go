package adapters

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/michalq/imgw/internal/climate"
)

func ProviderSqlite(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}

type SqliteMeasurementsRepository struct {
	db *sql.DB
}

func NewSqliteMeasurementsRepository(db *sql.DB) *SqliteMeasurementsRepository {
	return &SqliteMeasurementsRepository{db}
}

func (s *SqliteMeasurementsRepository) Setup(ctx context.Context) error {
	setupQuery := `
drop table if exists measurements;
create table measurements
(
    stationId text,
    day       text,
    y         integer,
    m         integer,
    d         integer,
    t         float,
    tmax      float,
    tmin      float
);

drop table if exists stations;
create table stations
(
    id            text,
    name          text,
    oldestMeasure text,
    newestMeasure text
);
`
	if _, err := s.db.Exec(setupQuery); err != nil {
		return err
	}
	return nil
}

func (s *SqliteMeasurementsRepository) AddMeasurements(
	ctx context.Context, measurements []climate.Daily,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	stmt, err := tx.Prepare(`insert into measurements (stationId, day, y, m, d, t, tmax, tmin) values (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	for _, daily := range measurements {
		_, err := stmt.Exec(daily.StationId, daily.Date(), daily.Year, daily.Month, daily.Day, daily.AvgTemp, daily.MaxTemp, daily.MinTemp)
		if err != nil {
			return fmt.Errorf("failed to insert measurement: %w", err)
		}
	}
	if err := stmt.Close(); err != nil {
		return fmt.Errorf("failed to close statement: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (s *SqliteMeasurementsRepository) AddStations(
	ctx context.Context, stations []climate.Station,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	stmt, err := tx.Prepare(`insert into stations (id,name,oldestMeasure,newestMeasure) values (?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	for _, station := range stations {
		_, err := stmt.Exec(station.Id, station.Name, station.OldestMeasurement, station.NewestMeasurement)
		if err != nil {
			return fmt.Errorf("failed to insert stations: %w", err)
		}
	}
	if err := stmt.Close(); err != nil {
		return fmt.Errorf("failed to close statement: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
