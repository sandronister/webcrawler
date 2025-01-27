package scrapping

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/ports"
	"github.com/sandronister/webcrawler/pkg/broker_cache/redis/types"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

type Model struct {
	server     ports.Iserver
	broker     types.IBroker
	logger     typelogger.ILogger
	reader     ports.IReader
	enviroment *config.Enviroment
}

func NewScracppingService(server ports.Iserver, broker types.IBroker, logger typelogger.ILogger, reader ports.IReader, enviroment *config.Enviroment) *Model {
	return &Model{
		server:     server,
		broker:     broker,
		logger:     logger,
		reader:     reader,
		enviroment: enviroment,
	}
}
