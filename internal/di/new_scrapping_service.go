package di

import (
	"github.com/sandronister/webcrawler/config"
	scrapping "github.com/sandronister/webcrawler/internal/service/scrappping"
	"github.com/sandronister/webcrawler/pkg/logger/types"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/factory"
)

func NewScracppingService(enviroment *config.Enviroment, logger types.ILogger) (*scrapping.Model, error) {

	broker, err := factory.GetBroker(enviroment.BrokerHost, enviroment.BrokerPort)

	if err != nil {
		return nil, err
	}

	cacher, err := factory.GetCacher(enviroment.BrokerHost, enviroment.BrokerPort, enviroment.BrokerDB)

	if err != nil {
		return nil, err
	}

	server := NewServer(broker, cacher, logger, enviroment)
	reader, err := NewReader(logger, broker, cacher, enviroment)

	if err != nil {
		return nil, err
	}

	scrapping := scrapping.NewScracppingService(server, broker, logger, reader, enviroment)

	return scrapping, nil
}
