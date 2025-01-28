package page

import (
	"github.com/sandronister/webcrawler/config"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

type Model struct {
	broker types.IBroker
	cacher types.ICacher
	env    *config.Enviroment
	logger typelogger.ILogger
}

func NewPageUsecase(broker types.IBroker, cacher types.ICacher, logger typelogger.ILogger, env *config.Enviroment) *Model {
	return &Model{
		broker: broker,
		cacher: cacher,
		env:    env,
		logger: logger,
	}
}
