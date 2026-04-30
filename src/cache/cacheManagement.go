package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func ConectarRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis_proj_bd:6379",
		Password: "",
		DB:       0,
	})
	err := RedisClient.Ping(Ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("Não foi possível conectar ao Redis: %v", err))
	}
	fmt.Println("Redis conectado.")
}
