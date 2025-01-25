package usecase

import (
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
		Topic: "page",
		Value: []byte(url),
	}

	return u.broker.Producer(message)
}
