package types

import "time"

type iserverMemData interface {
	Ping() error
}

type IBroker interface {
	iserverMemData
	ListenToQueue(conf *ConfigBroker, message chan<- Message) error
	Publish(message *Message) error
}

type ICacher interface {
	iserverMemData
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Exists(key string) bool
	Delete(key string) error
	Keys() ([]string, error)
}

type ISystemMemoryData interface {
	IBroker
	ICacher
}
