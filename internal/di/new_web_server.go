package di

import (
	"github.com/sandronister/webcrawler/config"
	echoserver "github.com/sandronister/webcrawler/internal/infra/web/echo_server"
	"github.com/sandronister/webcrawler/internal/ports/iserver"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

func NewServer(broker types.IBroker, cacher types.ICacher, logger typelogger.ILogger, env *config.Enviroment) iserver.Type {
	server := echoserver.NewServer(env.WebPort)
	handler := NewPageHandler(broker, cacher, logger, env)
	server.AddPageHandler(handler)
	return server
}
