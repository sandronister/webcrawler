package crawler

import (
	"github.com/sandronister/webcrawler/pkg/logger/types"
)

type Crawler struct {
	links  []string
	logger types.ILogger
}

func NewCrawler(logger types.ILogger) *Crawler {
	return &Crawler{
		links:  []string{},
		logger: logger,
	}
}
