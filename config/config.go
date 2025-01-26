package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Enviroment struct {
	WebPort     string `envconfig:"WEB_PORT"`
	BrokerHost  string `envconfig:"BROKER_HOST"`
	BrokerPort  string `envconfig:"BROKER_PORT"`
	BrokerTopic string `envconfig:"BROKER_TOPIC"`
	LogPattern  string `envconfig:"LOG_PATTERN"`
	BrokerKind  string `envconfig:"BROKER_KIND"`
}

func LoadEnviroment() (*Enviroment, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var env Enviroment
	err = envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}
