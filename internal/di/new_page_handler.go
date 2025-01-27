package di

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/infra/handler/page"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewPageHandler(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) *page.Model {
	return page.NewPageHandler(NewPageUseCase(broker, logger, env))
}
