package config

import (
	"github.com/spf13/viper"
)

type Enviroment struct {
	WebPort    string `mapstructure:"WEB_PORT"`
	BrokerPort int    `mapstructure:"BROKER_PORT"`
	BrokerHost string `mapstructure:"BROKER_HOST"`
	BrokerKind string `mapstructure:"BROKER_KIND"`
}

func LoadEnviroment(path string) (*Enviroment, error) {
	var enviroment *Enviroment

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&enviroment); err != nil {
		return nil, err
	}

	return enviroment, nil
}
