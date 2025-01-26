package scrapping

import (
	"encoding/json"

	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/internal/dto"
)

func (m *Model) WebServer() {
	m.server.Start()
}

func (m *Model) ListenToQueue(message chan<- types.Message) {
	configBroker := &types.ConfigBroker{
		Topic: []string{m.enviroment.BrokerTopic},
	}

	err := m.broker.ListenToQueue(configBroker, message)

	if err != nil {
		m.logger.Error("Error to consume message %s", err.Error())
	}
}

func (m *Model) ReadMessage(message <-chan types.Message) {
	var dto *dto.PageDTO

	for msg := range message {
		err := json.Unmarshal(msg.Value, dto)
		if err != nil {
			m.logger.Error("Error to unmarshal message %s", err.Error())
		}
		m.reader.Read(dto)
	}
}
