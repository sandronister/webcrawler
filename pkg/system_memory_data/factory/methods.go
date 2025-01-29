package factory

import (
	"fmt"

	"github.com/sandronister/webcrawler/pkg/system_memory_data/redis"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

func GetBroker(host string, port int) (types.IBroker, error) {
	if host == "" || port == 0 {
		return nil, fmt.Errorf("invalid Redis configuration: host and port must be specified")
	}

	broker := redis.NewBroker(host, port)

	if err := broker.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: please check if the Redis server is running and accessible")
	}

	return broker, nil
}

func GetCacher(host string, port, db int) (types.ICacher, error) {
	cacher := redis.NewCacher(host, port, db)

	if err := cacher.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: please check if the Redis server is running and accessible")
	}

	return cacher, nil
}
