package icrawler

type Type interface {
	Crawl(url string) (string, error)
}
