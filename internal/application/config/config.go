package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	AppEnv        string `env:"APP_ENV"`
	ServerAddress string `env:"ADDRESS"`
	AppKey        string `env:"APP_KEY"`
	DB            struct {
		DSN    string `env:"DATABASE_DSN"`
		Driver string `env:"DATABASE_DRIVER" envDefault:"pgx"`
	}
	Mail struct {
		UserName      string `env:"MAIL_USER_NAME"`
		UserPassword  string `env:"MAIL_USER_PASSWORD"`
		FromAddress   string `env:"MAIL_FROM_ADDRESS"`
		Host          string `env:"MAIL_HOST"`
		Port          string `env:"MAIL_PORT"`
		SendPerSecond int    `env:"MAIL_SEND_PER_SECOND"`
		QueueSize int    `env:"MAIL_QUEUE_SIZE" envDefault:"100"`
	}
	Log struct {
		Directory string `env:"LOG_DIRECTORY" envDefault:"storage/logs/"`
		MaxSize int `env:"LOG_MAX_SIZE" envDefault:"1"`
		MaxBackups int `env:"LOG_MAX_BACKUPS" envDefault:"1"`
		MaxAge  int `env:"LOG_MAX_AGE" envDefault:"1"`
		Compress bool `env:"LOG_Compress" envDefault:"true"`
	}
}

func (c Config) Error() string {
	panic("implement me")
}

var config Config
var err error
var once sync.Once

func GetConfig() (*Config, error) {
	once.Do(func() {
		config, err = NewConfigBuilder().WithAllFromEnv().Build()
	})
	return &config, err
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
