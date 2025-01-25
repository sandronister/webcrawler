package redisreader

import (
	"fmt"
	"time"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (r *Model) Read(url string) {

	link := make(chan string)

	content, err := r.ReadContent(url)

	if err != nil {
		return
	}

	for range 5 {
		go r.SendLink(link)
	}

	r.parser.ExtractLinks(content, link)

}

func (r *Model) ReadContent(url string) (string, error) {
	time.Sleep(30 * time.Second)

	content, err := r.crawler.Crawl(url)

	if err != nil {
		r.log.Error("Reader", fmt.Sprintf("Error crawling url %s: %s", url, err.Error()))
		return "", err
	}

	err = r.SaveContent(url, content)

	if err != nil {
		return "", err
	}

	return content, nil
}

func (r *Model) SaveContent(url, content string) error {
	err := r.repository.Insert(url, content)

	if err != nil {
		r.log.Error("Reader", fmt.Sprintf("Error inserting content: %s from url %s", err.Error(), url))
		return err
	}

	return nil
}

func (r *Model) SendLink(link <-chan string) {
	for l := range link {
		msg := &types.Message{
			Topic: "pages",
			Value: []byte(l),
		}
		r.broker.Producer(msg)
	}
}
