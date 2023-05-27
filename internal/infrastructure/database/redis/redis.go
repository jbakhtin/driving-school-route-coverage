package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jbakhtin/driving-school-route-coverage/internal/application/config"
)

type Redis struct {
	redis.Client
}

func New(config config.Config) (*Redis, error) {
	var client Redis
	client.Client = *redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}


	return &client, nil
}
