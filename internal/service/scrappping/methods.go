package scrapping

import (
	"fmt"

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
		fmt.Println("Error to consume message")
		m.logger.Error("Error to consume message %s", err.Error())
	}
}

func (m *Model) ReadMessage(message <-chan types.Message) {
	for msg := range message {
		fmt.Println("Message: ", string(msg.Value))
		m.reader.Read(string(msg.Value))
	}
}
