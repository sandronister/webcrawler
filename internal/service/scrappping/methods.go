package scrapping

import (
	"encoding/json"

	"github.com/sandronister/webcrawler/internal/dto"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
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

	for msg := range message {

		var dtoPage dto.PageDTO

		err := json.Unmarshal(msg.Value, &dtoPage)

		if err != nil {
			m.logger.Error("Error to unmarshal message %s", err.Error())
		}

		m.reader.Read(&dtoPage)
	}
}
