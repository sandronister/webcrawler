package redisreader

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/internal/ports"
)

type Model struct {
	crawler    ports.ICrawler
	parser     ports.IParser
	repository ports.IRepository
	log        ports.ILog
	broker     types.IBroker
}

func NewReader(crawler ports.ICrawler, parser ports.IParser, repository ports.IRepository, log ports.ILog, broker types.IBroker) *Model {
	return &Model{
		crawler:    crawler,
		parser:     parser,
		repository: repository,
		broker:     broker,
		log:        log,
	}
}
