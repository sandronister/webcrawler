package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sandronister/webcrawler/pkg/broker_cache/types"
)

func (m *Model) ListenToQueue(config *types.ConfigBroker, message chan<- types.Message) error {
	if config == nil {
		return types.ErrInvalidConfig
	}

	if config.Topic == nil {
		return types.ErrInvalidConfig
	}

	for {
		res, err := m.client.BLPop(context.Background(), 0*time.Second, config.Topic...).Result()
		if err != nil {
			fmt.Println("Erro ao ler item da fila:", err)
			continue
		}

		var tmp types.Message

		json.Unmarshal([]byte(res[1]), &tmp)

		message <- tmp

	}
}

func (m *Model) Publish(message *types.Message) error {
	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return m.client.LPush(context.Background(), message.Topic, messageByte).Err()
}

func (m *Model) Ping() error {
	return m.client.Ping(context.Background()).Err()
}
