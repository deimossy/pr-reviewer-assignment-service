package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"go.uber.org/zap"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT" envDefault:"8080"`

	ServerReadTimeout       time.Duration `env:"SERVER_HTTP_READ_TIMEOUT" envDefault:"5s"`
	ServerWriteTimeout      time.Duration `env:"SERVER_HTTP_WRITE_TIMEOUT" envDefault:"10s"`
	ServerIdleTimeout       time.Duration `env:"SERVER_HTTP_IDLE_TIMEOUT" envDefault:"120s"`
	ServerReadHeaderTimeout time.Duration `env:"SERVER_HTTP_READ_HEADER_TIMEOUT" envDefault:"5s"`

	PgDSN             string        `env:"PG_DSN" envDefault:"postgresql://app:app@postgres:5432/app?sslmode=disable"`
	PgMaxOpenConns    int           `env:"PG_MAX_OPEN_CONNS" envDefault:"10"`
	PgMaxIdleConns    int           `env:"PG_MAX_IDLE_CONNS" envDefault:"5"`
	PgConnMaxLifetime time.Duration `env:"PG_CONN_MAX_LIFETIME" envDefault:"1h"`
}

func LoadConfig(log *zap.Logger) (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Error("failed to load config", zap.Error(err))
		return nil, err
	}

	log.Info("config loaded", zap.String("server_port", cfg.ServerPort))

	return cfg, nil
}
