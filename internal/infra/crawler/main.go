package crawler

import (
	"io"
	"net/http"

	"github.com/sandronister/webcrawler/internal/ports"
)

type Crawler struct {
	links   []string
	visited []string
	logger  ports.ILog
}

func NewCrawler(logger ports.ILog) *Crawler {
	return &Crawler{
		links:   []string{},
		visited: []string{},
		logger:  logger,
	}
}

func (c *Crawler) isVisited(url string) bool {
	for _, v := range c.visited {
		if v == url {
			return true
		}
	}

	c.visited = append(c.visited, url)
	return false
}

func (c *Crawler) Crawl(curl <-chan string, cContent chan<- string, cHTML chan<- string) {
	for url := range curl {
		if c.isVisited(url) {
			return
		}

		resp, err := http.Get(url)

		if err != nil {
			c.logger.Log("Crawler", err.Error())
		}

		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)

		if err != nil {
			c.logger.Log("Crawler", err.Error())
		}

		cHTML <- string(content)
		cContent <- string(content)
	}
}
