package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Logger struct {
		Level string
		File  string
	}

	HTTP struct {
		Port string
	}
}

func New(confFile string) (*Config, error) {
	if confFile == "" {
		return nil, fmt.Errorf("confFile can't be empty")
	}

	viper.SetConfigFile(confFile)

	config := &Config{}

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
