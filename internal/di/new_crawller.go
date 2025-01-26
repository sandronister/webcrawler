package di

import (
	"github.com/sandronister/webcrawler/internal/infra/crawler"
	"github.com/sandronister/webcrawler/internal/ports"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func newCrawler(logger typelogger.ILogger) ports.ICrawler {
	return crawler.NewCrawler(logger)
}
