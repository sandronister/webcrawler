package di

import (
	"github.com/sandronister/webcrawler/config"
	redisreader "github.com/sandronister/webcrawler/internal/infra/reader/redis_reader"
	"github.com/sandronister/webcrawler/internal/ports/ireader"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

func NewReader(logger typelogger.ILogger, broker types.IBroker, env *config.Enviroment) (ireader.Type, error) {
	parser := newParser()
	repository, err := newRepository(parser, logger, env)

	if err != nil {
		return nil, err
	}
	return redisreader.NewReader(newCrawler(logger), parser, repository, logger, broker, env), nil
}
