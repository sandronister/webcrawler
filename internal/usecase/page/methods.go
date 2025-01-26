package page

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (u *Model) GetPage(url string) error {
	message := &types.Message{
		Topic: u.env.BrokerTopic,
		Value: []byte(url),
	}

	return u.broker.Publish(message)
}
