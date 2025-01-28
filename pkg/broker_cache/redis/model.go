package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
)

type Model struct {
	client *redis.Client
}

func NewBroker(server string, port int) (types.IBroker, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", server, port),
	})

	if _, err := client.Ping(client.Context()).Result(); err != nil {
		return nil, err
	}

	return &Model{
		client: client,
	}, nil
}
