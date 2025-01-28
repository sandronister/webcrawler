package redis

import "time"

func (m *Model) Delete(key string) error {
	return m.client.Del(m.client.Context(), key).Err()
}

func (m *Model) Get(key string) (string, error) {
	return m.client.Get(m.client.Context(), key).Result()
}

func (m *Model) Set(key string, value string, expiration time.Duration) error {
	return m.client.Set(m.client.Context(), key, value, expiration).Err()
}

func (m *Model) Exists(key string) bool {
	exists, err := m.client.Exists(m.client.Context(), key).Result()
	if err != nil {
		return false
	}
	return exists > 0
}

func (m *Model) Keys() ([]string, error) {
	return m.client.Keys(m.client.Context(), "*").Result()
}
