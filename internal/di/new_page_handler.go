package di

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/infra/handler/page"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewPageHandler(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) *page.Model {
	return page.NewPageHandler(NewPageUseCase(broker, logger, env))
}
