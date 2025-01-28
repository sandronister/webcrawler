package redisreader

import (
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/ports/icrawler"
	"github.com/sandronister/webcrawler/internal/ports/iparser"
	"github.com/sandronister/webcrawler/internal/ports/irepository"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

type Model struct {
	crawler    icrawler.Type
	parser     iparser.Type
	repository irepository.Type
	log        typelogger.ILogger
	broker     types.IBroker
	cacher     types.ICacher
	env        *config.Enviroment
}

func NewReader(crawler icrawler.Type, parser iparser.Type, repository irepository.Type, log typelogger.ILogger, broker types.IBroker, cacher types.ICacher, env *config.Enviroment) *Model {
	return &Model{
		crawler:    crawler,
		parser:     parser,
		repository: repository,
		broker:     broker,
		cacher:     cacher,
		log:        log,
		env:        env,
	}
}
