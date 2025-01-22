package main

import (
	"fmt"

	"github.com/sandronister/webcrawler/internal/web/crawler"
)

func main() {
	url := "https://comprecar.com.br"
	fmt.Printf("Fetching URL: %s\n", url)

	crawler := crawler.NewCrawler()
	links := crawler.Crawl(url)

	fmt.Println("Links encontrados:")

	for _, link := range links {
		fmt.Println(link)
	}

}
