package redisreader

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sandronister/webcrawler/internal/dto"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
)

func (r *Model) Read(message *dto.PageDTO) {

	fmt.Printf("Reading url %s\n", message.URL)
	link := make(chan dto.PageDTO)

	if r.env.TimeSleep > 0 {
		time.Sleep(time.Duration(r.env.TimeSleep) * time.Second)
	}

	content, err := r.ReadContent(message.URL)

	if err != nil {
		return
	}

	for range 10 {
		go r.SendLink(link)
	}

	r.parser.ExtractLinks(content, message.Filter, link)

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

func (r *Model) SendLink(link <-chan dto.PageDTO) {
	for l := range link {
		content, err := json.Marshal(l)
		if err != nil {
			r.log.Error("Reader", fmt.Sprintf("Error marshalling link: %s", err.Error()))
			fmt.Println("DEU RUIM")
		}

		msg := &types.Message{
			Topic: r.env.BrokerTopic,
			Value: content,
		}
		r.broker.Publish(msg)
	}
}
