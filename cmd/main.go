package main

import (
	"fmt"
	"sync"

	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/di"
	"github.com/sandronister/webcrawler/internal/ports"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func webServer(env *config.Enviroment, broker types.IBroker) {
	server := di.NewServer(env.WebPort, broker)
	server.Start()

}

func consumerBroker(broker types.IBroker, message chan<- types.Message, logger typelogger.ILogger) {
	configBroker := &types.ConfigBroker{
		Topic: []string{"page"},
	}

	err := broker.Consumer(configBroker, message)

	if err != nil {
		fmt.Println("Error to consume message")
		logger.Error("Error to consume message %s", err)
	}
}

func getLogger(env *config.Enviroment) typelogger.ILogger {
	logger, err := di.NewLogger(env.LogPattern)

	if err != nil {
		panic(err)
	}

	return logger
}

func readMessage(message <-chan types.Message, reader ports.IReader) {
	for m := range message {
		fmt.Println("Message: ", string(m.Value))
		reader.Read(string(m.Value))
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	message := make(chan types.Message)
	env, err := config.LoadEnviroment(".")
	if err != nil {
		panic(err)
	}

	logger := getLogger(env)

	broker := factory.NewBroker(env.BrokerKind, env.BrokerHost, env.BrokerPort)
	reader := di.NewReader(logger, broker)

	for range 8 {
		go readMessage(message, reader)
	}

	go consumerBroker(broker, message, logger)
	go webServer(env, broker)

	wg.Wait()
}
