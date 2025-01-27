package redisreader

import (
	"fmt"
	"time"

	"github.com/sandronister/webcrawler/pkg/broker_cache/redis/types"
)

func (r *Model) Read(url string) {

	fmt.Printf("Reading url %s\n", url)
	link := make(chan string)

	time.Sleep(10 * time.Second)

	content, err := r.ReadContent(url)

	if err != nil {
		return
	}

	for range 10 {
		go r.SendLink(link)
	}

	r.parser.ExtractLinks(content, link)

}

func (r *Model) ReadContent(url string) (string, error) {
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
			Topic: r.env.BrokerTopic,
			Value: []byte(l),
		}
		r.broker.Publish(msg)
	}
}
