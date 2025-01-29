package regexparser

import (
	"github.com/sandronister/webcrawler/internal/ports/ifilter"
)

type Model struct {
	filter ifilter.Type
}

func NewRegexParser(filter ifilter.Type) *Model {
	return &Model{
		filter: filter,
	}
}
