package crawler

type Climate struct{}

func NewClimate() *Climate {
	return &Climate{}
}

func (*Climate) Packages() []Package {
	return append(append(
		// Between 1951-2000 there are yearly packages, ex. 1951_k.zip
		yearly(1951, 2000),
		// 2000 up packages are per month, eg. 2001_01_k.zip
		monthly(2001, 2024)...,
	),
		// 2025 is not full, so each available month is manually added.
		formatMonthly(2025, 1),
		formatMonthly(2025, 2),
		formatMonthly(2025, 3),
		formatMonthly(2025, 4),
		formatMonthly(2025, 5),
		formatMonthly(2025, 6),
		formatMonthly(2025, 7),
		formatMonthly(2025, 8),
		formatMonthly(2025, 9),
		formatMonthly(2025, 10),
	)
}
