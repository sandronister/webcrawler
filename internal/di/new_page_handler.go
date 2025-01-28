package di

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/infra/handler/page"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

func NewPageHandler(broker types.IBroker, cacher types.ICacher, logger typelogger.ILogger, env *config.Enviroment) *page.Model {
	return page.NewPageHandler(NewPageUseCase(broker, cacher, logger, env))
}
