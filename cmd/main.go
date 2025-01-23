package main

import (
	"github.com/sandronister/webcrawler/internal/di"
	"github.com/sandronister/webcrawler/internal/infra/logCrawller"
)

func main() {
	url := "https://apinfo.com"

	logger := logCrawller.NewLog()
	logger.Info("main", "Starting webcrawler")

	reader := di.NewReader(logger)

	reader.Read(url)
}
