package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Enviroment struct {
	WebPort        string `envconfig:"WEB_PORT"`
	BrokerHost     string `envconfig:"BROKER_HOST"`
	BrokerPort     int    `envconfig:"BROKER_PORT"`
	BrokerTopic    string `envconfig:"BROKER_TOPIC"`
	BrokerDB       int    `envconfig:"BROKER_DB"`
	LogPattern     string `envconfig:"LOG_PATTERN"`
	BrokerKind     string `envconfig:"BROKER_KIND"`
	TimeSleep      int    `envconfig:"TIME_SLEEP"`
	RepositoryKind string `envconfig:"REPOSITORY_KIND"`
	RepositoryFile string `envconfig:"REPOSITORY_FILE"`
}

func LoadEnviroment() (*Enviroment, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	var env Enviroment
	err = envconfig.Process("", &env)
	if err != nil {
		return nil, fmt.Errorf("error loading enviroment variables: %w", err)
	}
	return &env, nil
}
