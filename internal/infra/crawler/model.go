package crawler

import (
	"github.com/sandronister/webcrawler/internal/ports"
)

type Crawler struct {
	links   []string
	visited []string
	logger  ports.ILog
}

func NewCrawler(logger ports.ILog) *Crawler {
	return &Crawler{
		links:   []string{},
		visited: []string{},
		logger:  logger,
	}
}
