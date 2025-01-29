package di

import (
	regexparser "github.com/sandronister/webcrawler/pkg/parser_html/regex_parser"
	"github.com/sandronister/webcrawler/pkg/parser_html/types"
)

func newParser() types.IPort {
	return regexparser.NewRegexParser(newFilter())
}
