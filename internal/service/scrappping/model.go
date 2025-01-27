package scrapping

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/ports/ireader"
	"github.com/sandronister/webcrawler/internal/ports/iserver"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

type Model struct {
	server     iserver.Type
	broker     types.IBroker
	logger     typelogger.ILogger
	reader     ireader.Type
	enviroment *config.Enviroment
}

func NewScracppingService(server iserver.Type, broker types.IBroker, logger typelogger.ILogger, reader ireader.Type, enviroment *config.Enviroment) *Model {
	return &Model{
		server:     server,
		broker:     broker,
		logger:     logger,
		reader:     reader,
		enviroment: enviroment,
	}
}
