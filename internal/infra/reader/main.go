package reader

import (
	"fmt"

	"github.com/sandronister/webcrawler/internal/ports"
)

type Reader struct {
	crawler    ports.ICrawler
	parser     ports.IParser
	repository ports.IRepository
}

func NewReader(crawler ports.ICrawler, parser ports.IParser, repository ports.IRepository) *Reader {
	return &Reader{
		crawler:    crawler,
		parser:     parser,
		repository: repository,
	}
}

func (r *Reader) Load(cUrl <-chan string) {
	for url := range cUrl {
		_, err := r.crawler.Crawl(url)
		if err != nil {
			panic(err)
		}
		fmt.Println("url: ", url)
	}
}

func (r *Reader) Read(url string) {

	links := make(chan string)

	content, err := r.crawler.Crawl(url)

	if err != nil {
		panic(err)
	}

	for range 9 {
		go r.Load(links)
	}

	r.parser.ExtractLinks(content, links)

}
