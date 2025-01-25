package di

import (
	"github.com/sandronister/webcrawler/internal/infra/crawler"
	"github.com/sandronister/webcrawler/internal/infra/parser/html"
	"github.com/sandronister/webcrawler/internal/infra/reader"
	"github.com/sandronister/webcrawler/internal/infra/repository/file"
	"github.com/sandronister/webcrawler/internal/infra/system"
	"github.com/sandronister/webcrawler/internal/ports"
)

func NewReader(logger ports.ILog) *reader.Reader {
	parser := newParser()
	return reader.NewReader(newCrawler(logger), parser, newRepository(parser), logger)

}

func newCrawler(logger ports.ILog) ports.ICrawler {
	return crawler.NewCrawler(logger)
}

func newParser() ports.IParser {
	return html.NewHtmlParser()
}

func newRepository(parser ports.IParser) ports.IRepository {
	system := system.NewOS()
	return file.NewFileRepository("output", parser, system)
}
