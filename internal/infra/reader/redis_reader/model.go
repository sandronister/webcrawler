package redisreader

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/ports"
	"github.com/sandronister/webcrawler/pkg/broker_cache/redis/types"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

type Model struct {
	crawler    ports.ICrawler
	parser     ports.IParser
	repository ports.IRepository
	log        typelogger.ILogger
	broker     types.IBroker
	env        *config.Enviroment
}

func NewReader(crawler ports.ICrawler, parser ports.IParser, repository ports.IRepository, log typelogger.ILogger, broker types.IBroker, env *config.Enviroment) *Model {
	return &Model{
		crawler:    crawler,
		parser:     parser,
		repository: repository,
		broker:     broker,
		log:        log,
		env:        env,
	}
}
