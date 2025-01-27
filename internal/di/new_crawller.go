package di

import (
	"github.com/sandronister/webcrawler/internal/infra/crawler"
	"github.com/sandronister/webcrawler/internal/ports/icrawler"
	"github.com/sandronister/webcrawler/pkg/logger/types"
)

func newCrawler(logger types.ILogger) icrawler.Type {
	return crawler.NewCrawler(logger)
}
