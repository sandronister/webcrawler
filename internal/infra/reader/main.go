package reader

import "github.com/sandronister/webcrawler/internal/ports"

type Reader struct {
	crawler    ports.ICrawler
	parser     ports.IParser
	repository ports.IRepository
	cUrl       chan string
	cContent   chan string
	cHTML      chan string
}

func NewReader(crawler ports.ICrawler, parser ports.IParser, repository ports.IRepository) *Reader {
	return &Reader{
		crawler:    crawler,
		parser:     parser,
		repository: repository,
		cUrl:       make(chan string),
		cContent:   make(chan string),
		cHTML:      make(chan string),
	}
}

func (r *Reader) Read(url string) {
	r.cUrl <- url

	for range 5 {
		go r.repository.Insert(r.cContent)
	}

	for range 5 {
		go r.parser.ExtractLinks(r.cHTML, r.cUrl)
	}

	for range 5 {
		go r.crawler.Crawl(r.cUrl, r.cContent, r.cHTML)
	}
}
