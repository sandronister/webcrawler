package crawler

import (
	"io"
	"net/http"
)

func (c *Crawler) isVisited(url string) bool {
	for _, v := range c.visited {
		if v == url {
			return true
		}
	}

	c.visited = append(c.visited, url)
	return false
}

func (c *Crawler) Crawl(url string) (string, error) {
	if c.isVisited(url) {
		return "", nil
	}

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(content), nil

}
