package crawler

import (
	"io"
	"net/http"
	"regexp"
)

type Crawler struct {
	links []string
}

func NewCrawler() *Crawler {
	return &Crawler{
		links: []string{},
	}
}

func (c *Crawler) Crawl(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return c.ExtractLinks(string(content)), nil

}

func (c *Crawler) ExtractLinks(htmlContent string) []string {
	var links []string
	re := regexp.MustCompile(`href="(http[s]?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(htmlContent, -1)
	for _, match := range matches {
		if len(match) > 1 {
			links = append(links, match[1])
		}
	}
	return links
}
