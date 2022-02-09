package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Logger   `yaml:"logger"`
	HTTP     `yaml:"http"`
	Postgres `yaml:"postgres"`
}

type Logger struct {
	Level string `env-required:"true" yaml:"level"`
	File  string `env-required:"true" yaml:"file"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port"`
}

type Postgres struct {
	DSN string `env-required:"true" env:"PG_DSN"`
}

func New(confFile string) (*Config, error) {
	if confFile == "" {
		return nil, fmt.Errorf("confFile can't be empty")
	}

	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	config := &Config{}
	err := cleanenv.ReadConfig(confFile, config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
