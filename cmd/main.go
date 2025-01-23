package main

import (
	"fmt"

	"github.com/sandronister/webcrawler/internal/di"
	"github.com/sandronister/webcrawler/internal/infra/logger"
	"github.com/sandronister/webcrawler/internal/infra/system"
)

func main() {
	var url string
	fmt.Print("Informe a URL: ")
	fmt.Scanln(&url)
	fmt.Println("URL informada: ", url)

	system := system.NewOS()
	logger := logger.NewLog(system)
	logger.Info("main", "Starting webcrawler")

	reader := di.NewReader(logger)

	reader.Read(url)
}
