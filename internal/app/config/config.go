package config

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/creasty/defaults"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Configs struct {
	App      App
	Postgres Postgres
}

type App struct {
	Timeout time.Duration `env:"APP_TIMEOUT" default:"60s"`
	Port    string        `env:"APP_PORT"`
	Host    string        `env:"APP_HOST"`
	Env     string        `env:"APP_ENV"`
}

type Postgres struct {
	Host     string        `env:"POSTGRES_HOST"`
	Port     int           `env:"POSTGRES_PORT"`
	User     string        `env:"POSTGRES_USER"`
	Password string        `env:"POSTGRES_PASSWORD"`
	DBName   string        `env:"POSTGRES_DB_NAME"`
	SSLMode  string        `env:"POSTGRES_SSL_MODE"`
	Timeout  time.Duration `env:"POSTGRES_TIMEOUT" default:"20s"`
}

func New() (*Configs, error) {
	cfg := new(Configs)

	if err := godotenv.Load(); err != nil {
		log.Error("failed to load .env file")
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if err := defaults.Set(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
