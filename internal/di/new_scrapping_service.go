package di

import (
	"github.com/sandronister/webcrawler/config"
	scrapping "github.com/sandronister/webcrawler/internal/service/scrappping"
	"github.com/sandronister/webcrawler/pkg/broker_cache/redis"
)

func NewScracppingService(enviroment *config.Enviroment) (*scrapping.Model, error) {
	logger, err := NewLogger(enviroment.LogPattern)

	if err != nil {
		return nil, err
	}

	broker, err := redis.NewBroker("localhost", 6379)

	if err != nil {
		return nil, err
	}

	server := NewServer(broker, logger, enviroment)
	reader, err := NewReader(logger, broker, enviroment)

	if err != nil {
		return nil, err
	}

	scrapping := scrapping.NewScracppingService(server, broker, logger, reader, enviroment)

	return scrapping, nil
}
