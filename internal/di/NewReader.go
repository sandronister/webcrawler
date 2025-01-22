package di

import (
	"github.com/sandronister/webcrawler/internal/infra/crawler"
	"github.com/sandronister/webcrawler/internal/infra/parser"
	"github.com/sandronister/webcrawler/internal/infra/reader"
	"github.com/sandronister/webcrawler/internal/ports"
)

func NewReader(logger ports.ILog) *reader.Reader {
	return reader.NewReader(newCrawler(logger), newParser())

}

func newCrawler(logger ports.ILog) ports.ICrawler {
	return crawler.NewCrawler(logger)
}

func newParser() ports.IParser {
	return parser.NewHtmlParser()
}
