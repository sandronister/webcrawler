package di

import (
	"github.com/sandronister/webcrawler/internal/infra/parser/html"
	"github.com/sandronister/webcrawler/internal/ports/iparser"
)

func newParser() iparser.Type {
	return html.NewHtmlParser()
}
