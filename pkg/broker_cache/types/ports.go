package types

type IServerMemData interface {
	Ping() error
}

type IBroker interface {
	IServerMemData
	ListenToQueue(conf *ConfigBroker, message chan<- Message) error
	Publish(message *Message) error
}

type ICacher interface {
	IServerMemData
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Exists(key string) (bool, error)
	Delete(key string) error
	Keys() ([]string, error)
}

type ISystemMemoryData interface {
	IBroker
	ICacher
}
