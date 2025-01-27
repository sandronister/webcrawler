package di

import (
	"github.com/sandronister/webcrawler/config"
	pageusecase "github.com/sandronister/webcrawler/internal/usecase/page"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewPageUseCase(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) *pageusecase.Model {
	return pageusecase.NewPageUsecase(broker, logger, env)
}
