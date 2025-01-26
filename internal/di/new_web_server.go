package di

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	echoserver "github.com/sandronister/webcrawler/internal/infra/web/echo_server"
	"github.com/sandronister/webcrawler/internal/ports/iserver"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewServer(broker types.IBroker, logger typelogger.ILogger, env *config.Enviroment) iserver.Type {
	server := echoserver.NewServer(env.WebPort)
	handler := NewPageHandler(broker, logger, env)
	server.AddPageHandler(handler)
	return server
}
