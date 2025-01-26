package di

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	redisreader "github.com/sandronister/webcrawler/internal/infra/reader/redis_reader"
	"github.com/sandronister/webcrawler/internal/ports"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewReader(logger typelogger.ILogger, broker types.IBroker, env *config.Enviroment) (ports.IReader, error) {
	parser := newParser()
	repository, err := newRepository(parser, logger, env)

	if err != nil {
		return nil, err
	}
	return redisreader.NewReader(newCrawler(logger), parser, repository, logger, broker, env), nil
}
