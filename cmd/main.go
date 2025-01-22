package main

import (
	"github.com/sandronister/webcrawler/internal/di"
	"github.com/sandronister/webcrawler/internal/infra/logCrawller"
)

func main() {
	url := "https://www.sandronister.com.br"

	logger := logCrawller.NewLog()

	reader := di.NewReader(logger)

	reader.Read(url)
}
