package database

import (
	"be_dashboard/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
	Ctx   = context.Background()
)

func ConnectRedis() {

	cfg := config.AppConfig

	Redis = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	_, err := Redis.Ping(Ctx).Result()

	if err != nil {
		panic("Failed connect Redis : " + err.Error())
	}

	fmt.Println("Redis connected successfully! 🚀")
}
