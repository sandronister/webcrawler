package main

import (
	"github.com/sandronister/webcrawler/internal/di"
	"github.com/sandronister/webcrawler/internal/infra/logger"
	"github.com/sandronister/webcrawler/internal/infra/system"
)

func main() {
	url := "https://apinfo.com"

	system := system.NewOS()
	logger := logger.NewLog(system)
	logger.Info("main", "Starting webcrawler")

	reader := di.NewReader(logger)

	reader.Read(url)
}
