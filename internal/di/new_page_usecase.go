package di

import (
	"github.com/sandronister/webcrawler/config"
	pageusecase "github.com/sandronister/webcrawler/internal/usecase/page"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

func NewPageUseCase(broker types.IBroker, cacher types.ICacher, logger typelogger.ILogger, env *config.Enviroment) *pageusecase.Model {
	return pageusecase.NewPageUsecase(broker, cacher, logger, env)
}
