package reader

import "github.com/sandronister/webcrawler/internal/ports"

type Reader struct {
	crawler  ports.ICrawler
	parser   ports.IParser
	cUrl     chan string
	cContent chan string
}

func NewReader(crawler ports.ICrawler, parser ports.IParser) *Reader {
	return &Reader{
		crawler:  crawler,
		parser:   parser,
		cUrl:     make(chan string),
		cContent: make(chan string),
	}
}

func (r *Reader) Read(url string) {
	r.cUrl <- url

	for range 5 {
		go r.parser.ExtractLinks(r.cContent, r.cUrl)
	}

	for range 5 {
		go r.crawler.Crawl(r.cUrl, r.cContent)
	}
}
