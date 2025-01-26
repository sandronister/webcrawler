package di

import (
	"github.com/sandronister/webcrawler/internal/infra/parser/html"
	"github.com/sandronister/webcrawler/internal/ports"
)

func newParser() ports.IParser {
	return html.NewHtmlParser()
}
