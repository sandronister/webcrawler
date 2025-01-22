package reader

import "github.com/sandronister/webcrawler/internal/ports"

type Reader struct {
	crawler ports.ICrawler
}

func NewReader(crawler ports.ICrawler) *Reader {
	return &Reader{
		crawler: crawler,
	}
}

func (r *Reader) Read(url string) ([]string, error) {
	return r.crawler.Crawl(url)
}
