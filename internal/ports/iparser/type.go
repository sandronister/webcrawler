package iparser

import "github.com/sandronister/webcrawler/internal/dto"

type Type interface {
	GetTagContet(input, tag string) []string
	RemoveHTMLTags(input string) string
	ExtractLinks(content, urlFilter string, cUrl chan<- dto.PageDTO)
}
