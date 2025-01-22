package ports

type ICrawler interface {
	Crawl(url string) (string, error)
}
