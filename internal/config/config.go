package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PostgresConfig `yaml:"postgres`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
	
	// Postgresql config
	PostgresConfig struct {
		PostgresqlHost     string
		PostgresqlPort     string
		PostgresqlUser     string
		PostgresqlPassword string
		PostgresqlDbname   string
}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	
	err := cleanenv.ReadConfig("internal/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}