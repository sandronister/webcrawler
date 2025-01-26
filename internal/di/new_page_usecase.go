package di

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	pageusecase "github.com/sandronister/webcrawler/internal/usecase/page"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewPageUseCase(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) *pageusecase.Model {
	return pageusecase.NewPageUsecase(broker, logger, env)
}
