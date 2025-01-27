package page

import (
	"fmt"
	"net/http"

	"github.com/sandronister/webcrawler/pkg/broker_cache/redis/types"
)

func (u *Model) GetPage(url string) error {
	_, err := http.Get(url)

	if err != nil {
		u.logger.Error("Error getting the page", url)
		return fmt.Errorf("invalid url")
	}

	message := &types.Message{
		Topic: u.env.BrokerTopic,
		Value: []byte(url),
	}

	return u.broker.Publish(message)
}
