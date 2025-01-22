package main

import (
	"fmt"

	"github.com/sandronister/webcrawler/internal/infra/crawler"
)

func main() {
	url := "https://comprecar.com.br"
	fmt.Printf("Fetching URL: %s\n", url)

	crawler := crawler.NewCrawler()
	links, err := crawler.Crawl(url)

	if err != nil {
		fmt.Printf("Erro ao buscar URL: %s\n", err)
		return
	}

	fmt.Println("Links encontrados:")

	for _, link := range links {
		fmt.Println(link)
	}

}
