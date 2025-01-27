package di

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/infra/handler/page"
	echoserver "github.com/sandronister/webcrawler/internal/infra/web/echo_server"
	"github.com/sandronister/webcrawler/internal/ports"
	pageusecase "github.com/sandronister/webcrawler/internal/usecase/page"
	"github.com/sandronister/webcrawler/pkg/broker_cache/redis/types"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewServer(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) ports.Iserver {
	server := echoserver.NewServer(env.WebPort)
	handler := NewPageHandler(broker, logger, env)
	server.AddPageHandler(handler)
	return server
}

func NewPageHandler(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) *page.Model {
	return page.NewPageHandler(NewPageUseCase(broker, logger, env))
}

func NewPageUseCase(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) *pageusecase.Model {
	return pageusecase.NewPageUsecase(broker, logger, env)
}
