package crawler

import (
	"fmt"
	"strconv"
)

func formatYearly(y int) Package {
	return Package{
		Type:     PackageYearly,
		FileName: fmt.Sprintf("%d_k.zip", y),
		Year:     strconv.Itoa(y),
		Month:    "",
	}
}

func formatMonthly(y, m int) Package {
	return Package{
		Type:     PackageMonthly,
		FileName: fmt.Sprintf("%d_%02d_k.zip", y, m),
		Year:     strconv.Itoa(y),
		Month:    strconv.Itoa(m),
	}
}

func yearly(yearStart, yearEnd int) []Package {
	if yearEnd < yearStart {
		return make([]Package, 0)
	}
	all := make([]Package, 0, yearEnd-yearStart+1)
	for y := yearStart; y <= yearEnd; y++ {
		all = append(all, formatYearly(y))
	}
	return all
}

func monthly(yearStart, yearEnd int) []Package {
	if yearEnd < yearStart {
		return make([]Package, 0)
	}
	all := make([]Package, 0, yearEnd-yearStart+1)
	for y := yearStart; y <= yearEnd; y++ {
		for m := 1; m <= 12; m++ {
			all = append(all, formatMonthly(y, m))
		}
	}
	return all
}
