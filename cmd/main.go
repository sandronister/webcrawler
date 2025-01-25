package main

import (
	"sync"

	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/di"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	conf, err := config.LoadEnviroment(".")

	if err != nil {
		panic(err)
	}

	broker := factory.NewBroker(conf.BrokerKind, conf.BrokerHost, conf.BrokerPort)

	server := di.NewServer(conf.WebPort, broker)
	go server.Start()

	wg.Wait()
}
