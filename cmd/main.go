package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/di"
	"github.com/sandronister/webcrawler/internal/ports"
	typelogger "github.com/sandronister/webcrawler/pkg/logger/types"
)

func webServer(env *config.Enviroment, logger typelogger.ILogger, broker types.IBroker) {
	server := di.NewServer(env.WebPort, broker, logger, env)
	server.Start()

}

func consumerBroker(broker types.IBroker, message chan<- types.Message, logger typelogger.ILogger, env *config.Enviroment) {

	fmt.Printf("Consuming message from topic %s\n", env.BrokerTopic)
	configBroker := &types.ConfigBroker{
		Topic: []string{env.BrokerTopic},
	}

	err := broker.ListenToQueue(configBroker, message)

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
	env, err := config.LoadEnviroment()

	fmt.Printf("Env: %v\n", env.BrokerKind)

	fmt.Printf("Env: %v\n", os.Getenv("BROKER_KIND"))
	if err != nil {
		panic(err)
	}

	logger := getLogger(env)

	logger.Info("Starting webcrawler")

	broker, err := factory.GetBroker()
	if err != nil {
		logger.Error("Error to create broker %s", err)
		panic(err)
	}
	reader := di.NewReader(logger, broker, env)

	for range 8 {
		go readMessage(message, reader)
	}

	go consumerBroker(broker, message, logger, env)
	go webServer(env, logger, broker)

	fmt.Printf("Web server running on port %s\n", env.WebPort)
	logger.Info("Web server running on port %s", env.WebPort)

	wg.Wait()
}
