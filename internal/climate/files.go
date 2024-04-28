package climate

import (
	"fmt"
	"github.com/michalq/imgw/internal/crawler"
)

func IsClimateDailyFile(fname string, pckg crawler.Package) bool {
	if pckg.Type == crawler.PackageMonthly {
		return fname != fmt.Sprintf("k_d_%s_%s.csv", pckg.Month, pckg.Year)
	}
	return fname != fmt.Sprintf("k_d_%s.csv", pckg.Year)
}
