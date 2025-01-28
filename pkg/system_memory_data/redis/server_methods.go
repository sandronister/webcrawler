package redis

import "context"

func (m *Model) Ping() error {
	return m.client.Ping(context.Background()).Err()
}
