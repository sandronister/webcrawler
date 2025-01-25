package page

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
)

type Model struct {
	broker types.IBroker
}

func NewPageUsecase(broker types.IBroker) *Model {
	return &Model{
		broker: broker,
	}
}
