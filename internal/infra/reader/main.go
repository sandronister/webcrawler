package reader

import (
	"fmt"

	"github.com/sandronister/webcrawler/internal/ports"
)

type Reader struct {
	crawler    ports.ICrawler
	parser     ports.IParser
	repository ports.IRepository
	log        ports.ILog
}

func NewReader(crawler ports.ICrawler, parser ports.IParser, repository ports.IRepository, log ports.ILog) *Reader {
	return &Reader{
		crawler:    crawler,
		parser:     parser,
		repository: repository,
		log:        log,
	}
}

func (r *Reader) Load(cUrl <-chan string) {
	for url := range cUrl {
		fmt.Println(url)
		r.Read(url)
	}
}

func (r *Reader) Read(url string) {

	links := make(chan string)

	content, err := r.crawler.Crawl(url)

	if err != nil {
		r.log.Error("Reader", fmt.Sprintf("Error crawling url %s: %s", url, err.Error()))
		return
	}

	err = r.repository.Insert(url, content)

	if err != nil {
		r.log.Error("Reader", fmt.Sprintf("Error inserting content: %s from url %s", err.Error(), url))
	}

	for range 9 {
		go r.Load(links)
	}

	r.parser.ExtractLinks(content, links)

}
