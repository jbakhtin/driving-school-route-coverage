package config

import (
	"github.com/caarlos0/env/v6"
)


const (
	_serverAddress                    = "127.0.0.1:8080"
	_databaseDSN                = ""
	_databaseDriver             = "pgx"
)

type Config struct {
	ServerAddress  string `env:"ADDRESS"`
	DatabaseDSN    string `env:"DATABASE_DSN"`
	DatabaseDriver string `env:"DATABASE_DRIVER" envDefault:"pgx"`
}

type Builder struct {
	config Config
	err    error
}

func NewConfigBuilder() *Builder {
	return &Builder{
		Config{
			_serverAddress,
			_databaseDSN,
			_databaseDriver,
		},
		nil,
	}
}

func (cb *Builder) WithAllFromEnv() *Builder {
	err := env.Parse(&cb.config)
	if err != nil {
		cb.err = err
	}

	return cb
}

func (cb *Builder) Build() (Config, error) {
	return cb.config, cb.err
}