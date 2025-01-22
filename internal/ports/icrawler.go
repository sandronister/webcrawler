package ports

type ICrawler interface {
	Crawl(curl <-chan string, cContent chan<- string, cHTML chan<- string)
}
