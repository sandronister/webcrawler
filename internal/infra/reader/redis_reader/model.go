package redisreader

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/ports/icrawler"
	"github.com/sandronister/webcrawler/internal/ports/iparser"
	"github.com/sandronister/webcrawler/internal/ports/irepository"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

type Model struct {
	crawler    icrawler.Type
	parser     iparser.Type
	repository irepository.Type
	log        typelogger.ILogger
	broker     types.IBroker
	env        *config.Enviroment
}

func NewReader(crawler icrawler.Type, parser iparser.Type, repository irepository.Type, log typelogger.ILogger, broker types.IBroker, env *config.Enviroment) *Model {
	return &Model{
		crawler:    crawler,
		parser:     parser,
		repository: repository,
		broker:     broker,
		log:        log,
		env:        env,
	}
}
