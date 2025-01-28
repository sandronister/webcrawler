package main

import (
	"fmt"
	"sync"

	"github.com/sandronister/webcrawler/config"
	"github.com/sandronister/webcrawler/internal/di"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	message := make(chan types.Message)
	env, err := config.LoadEnviroment()

	if err != nil {
		fmt.Println("Error loading enviroment variables")
	}

	service, err := di.NewScracppingService(env)

	if err != nil {
		fmt.Printf("Error creating service: %s\n", err)
	}

	for range 8 {
		go service.ReadMessage(message)
	}

	go service.ListenToQueue(message)
	go service.WebServer()

	fmt.Printf("Web server running on port %s\n", env.WebPort)

	wg.Wait()
}
