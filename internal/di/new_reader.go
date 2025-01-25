package di

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/infra/crawler"
	"github.com/sandronister/webcrawler/internal/infra/parser/html"
	redisreader "github.com/sandronister/webcrawler/internal/infra/reader/redis_reader"
	"github.com/sandronister/webcrawler/internal/infra/repository/file"
	"github.com/sandronister/webcrawler/internal/infra/system"
	"github.com/sandronister/webcrawler/internal/ports"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func NewReader(logger typelogger.ILogger, broker types.IBroker, env *config.Enviroment) ports.IReader {
	parser := newParser()
	return redisreader.NewReader(newCrawler(logger), parser, newRepository(parser), logger, broker, env)

}

func newCrawler(logger typelogger.ILogger) ports.ICrawler {
	return crawler.NewCrawler(logger)
}

func newParser() ports.IParser {
	return html.NewHtmlParser()
}

func newRepository(parser ports.IParser) ports.IRepository {
	system := system.NewOS()
	return file.NewFileRepository("output", parser, system)
}
