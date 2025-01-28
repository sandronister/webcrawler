package factory

import (
	"github.com/sandronister/webcrawler/pkg/system_memory_data/redis"
	"github.com/sandronister/webcrawler/pkg/system_memory_data/types"
)

func GetBroker(host string, port int) (types.IBroker, error) {
	broker := redis.NewBroker(host, port)

	if err := broker.Ping(); err != nil {
		return nil, err
	}

	return broker, nil
}

func GetCacher(host string, port, db int) (types.ICacher, error) {
	cacher := redis.NewCacher(host, port, db)

	if err := cacher.Ping(); err != nil {
		return nil, err
	}

	return cacher, nil
}
