package crawler

type Climate struct{}

func NewClimate() *Climate {
	return &Climate{}
}

func (*Climate) Packages() []Package {
	return append(
		yearly(1951, 2000),
		monthly(2001, 2023)...,
	)
}
