package config

import (
	"github.com/caarlos0/env/v6"
	"sync"
)

type Config struct {
	ServerAddress  string `env:"ADDRESS"`
	AppKey  string `env:"APP_KEY"`
	DB             struct {
		DSN    string `env:"DATABASE_DSN"`
		Driver string `env:"DATABASE_DRIVER" envDefault:"pgx"`
	}
}

var config Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config, _ = NewConfigBuilder().WithAllFromEnv().Build()
	})
	return &config
}

type Builder struct {
	config Config
	err    error
}

func NewConfigBuilder() *Builder {
	return &Builder{
		Config{},
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
