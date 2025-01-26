package di

import (
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/webcrawler/config"
	scrapping "github.com/sandronister/webcrawler/internal/service/scrappping"
)

func NewScracppingService(enviroment *config.Enviroment) (*scrapping.Model, error) {
	logger, err := NewLogger(enviroment.LogPattern)

	if err != nil {
		return nil, err
	}

	broker, err := factory.GetBroker()

	if err != nil {
		return nil, err
	}

	server := NewServer(broker, logger, enviroment)
	reader := NewReader(logger, broker, enviroment)

	scrapping := scrapping.NewScracppingService(server, broker, logger, reader, enviroment)

	return scrapping, nil
}
