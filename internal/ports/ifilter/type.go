package ifilter

type Type interface {
	Filter(match []string, filter string) bool
}
