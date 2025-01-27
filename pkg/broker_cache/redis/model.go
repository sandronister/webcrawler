package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
)

type Model struct {
	client *redis.Client
}

func NewBroker(server string, port int) types.IBroker {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", server, port),
	})
	return &Model{
		client: client,
	}
}
