package cache

import (
	"context"
	"fmt"
	"todo_list/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitCache() {
	host := config.RedisHost
	port := config.RedisPort
	password := config.RedisPassword
	database := config.RedisDbName
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       database,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
