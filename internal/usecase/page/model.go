package page

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
)

type Model struct {
	broker types.IBroker
	env    *config.Enviroment
}

func NewPageUsecase(broker types.IBroker, env *config.Enviroment) *Model {
	return &Model{
		broker: broker,
		env:    env,
	}
}
