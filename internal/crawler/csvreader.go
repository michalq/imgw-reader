package crawler

type CsvReader[T any] interface {
	ParseLine([]string) T
}
