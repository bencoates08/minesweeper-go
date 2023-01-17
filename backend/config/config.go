package config

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	DB DatabaseConfig
}

type DatabaseConfig struct {
	User      string `env:"DB_USER" envDefault:"root"`
	Password  string `env:"DB_PASSWORD" envDefault:""`
	Container string `env:"DB_CONTAINER" envDefault:""`
	Name      string `env:"DB_NAME" envDefault:""`
}

func NewAppConfig() (AppConfig, error) {
	var cfg AppConfig

	err := env.Parse(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
