package page

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

type Model struct {
	broker types.IBroker
	env    *config.Enviroment
	logger typelogger.ILogger
}

func NewPageUsecase(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) *Model {
	return &Model{
		broker: broker,
		env:    env,
		logger: logger,
	}
}
