package crawler

import (
	"io"
	"net/http"
)

func (c *Crawler) Crawl(url string) (string, error) {

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
