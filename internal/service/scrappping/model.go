package scrapping

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/ports/ireader"
	"github.com/sandronister/webcrawler/internal/ports/iserver"
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
