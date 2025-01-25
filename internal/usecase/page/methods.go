package page

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
)

func (u *Model) GetPage(url string) error {
	message := &types.Message{
		Topic: "page",
		Value: []byte(url),
	}

	return u.broker.Producer(message)
}
