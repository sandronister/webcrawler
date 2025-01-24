package usecase

import (
	"time"

	"github.com/sandronister/go_broker/pkg/broker/types"
)

type PageUsecase struct {
	broker types.IBroker
}

func NewPageUsecase(broker types.IBroker) *PageUsecase {
	return &PageUsecase{
		broker: broker,
	}
}

func (u *PageUsecase) GetPage(url string) error {
	message := &types.Message{
		Topic:     "page",
		Value:     []byte(url),
		Timestamp: time.Now(),
	}

	return u.broker.Producer(message)

}
