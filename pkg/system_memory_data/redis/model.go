package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
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

func NewCacher(server string, port, db int) types.ICacher {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", server, port),
		DB:   db,
	})

	return &Model{
		client: client,
	}
}

func NewSystemMemoryData(server string, port, db int) types.ISystemMemoryData {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", server, port),
		DB:   db,
	})

	return &Model{
		client: client,
	}
}
