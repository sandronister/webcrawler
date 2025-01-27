package html

import (
	"github.com/sandronister/webcrawler/internal/ports/ifilter"
)

type Model struct {
	filter ifilter.Type
}

func NewHtmlParser(filter ifilter.Type) *Model {
	return &Model{
		filter: filter,
	}
}
