package synop

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/michalq/imgw/internal/climate"
)

func SynopScanner(debug bool, dir string) (climate.List, error) {
	climateParser := climate.NewDailyParser(climate.ParseClimateLine)
	measurements := make(climate.List, 0)
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		errs := make([]error, 0)
		if !d.IsDir() && filepath.Ext(d.Name()) == ".zip" {
			archive, err := zip.OpenReader(path)
			if err != nil {
				return fmt.Errorf("opening zip error %w", err)
			}
			for _, file := range archive.File {
				if file.Name[0:6] != "s_d_t_" {
					if debug {
						fmt.Println("Reading ", file.Name)
					}
					rc, err := file.Open()
					if err != nil {
						errs = append(errs, fmt.Errorf("could not open file %s (%w)", file.Name, err))
						continue
					}
					data, err := io.ReadAll(rc)
					if err != nil {
						errs = append(errs, fmt.Errorf("problem with reading file %s (%s)", file.Name, err))
						continue
					}
					r := csv.NewReader(bytes.NewReader(data))
					acqs, err := climateParser.ParseFromReader(r)
					if err != nil {
						errs = append(errs, fmt.Errorf("problem with parsing file %s (%s)", file.Name, err))
						continue
					}
					measurements = append(measurements, acqs...)
					if debug {
						fmt.Printf("Read %d daily measuremenets from package %s\n", len(acqs), file.Name)
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return measurements, nil
}
