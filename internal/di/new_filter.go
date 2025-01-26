package di

import (
	"github.com/sandronister/webcrawler/internal/infra/filter"
	"github.com/sandronister/webcrawler/internal/ports/ifilter"
)

func newFilter() ifilter.Type {
	return filter.NewURLFilter()
}
