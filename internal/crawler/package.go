package crawler

type PackageType string

const (
	PackageMonthly PackageType = "monthly"
	PackageYearly  PackageType = "yearly"
)

// Package represents single package of data.
type Package struct {
	Type     PackageType
	FileName string
	Year     string
	Month    string
}
