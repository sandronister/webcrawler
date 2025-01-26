package scrapping

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
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
		m.reader.Read(string(msg.Value))
	}
}
