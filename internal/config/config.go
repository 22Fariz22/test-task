package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		HTTP           `yaml:"http"`
		Log            `yaml:"logger"`
		PostgresConfig `yaml:"postgres"`
	}

	// HTTP -.
	HTTP struct {
		Address string `env-required:"true" yaml:"address" env:"HTTP_ADDRESS"`
		Port    string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Postgresql config
	PostgresConfig struct {
		PostgresqlHost     string `env-required:"true" yaml:"pg_host" env:"PG_HOST"`
		PostgresqlPort     string `env-required:"true" yaml:"pg_port" env:"PG_PORT"`
		PostgresqlUser     string `env-required:"true" yaml:"pg_user" env:"PG_USER"`
		PostgresqlPassword string `env-required:"true" yaml:"pg_password" env:"PG_PASSWORD"`
		PostgresqlDbname   string `env-required:"true" yaml:"pg_dbname" env:"PG_DBNAME"`
		PostgresqlSSLMode  bool
		PgDriver           string
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
