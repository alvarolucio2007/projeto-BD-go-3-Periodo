// Package cache faz a conexão e manipulação de cache.
package cache

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func ConectarRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "cache:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	err := rdb.Ping(Ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("Não foi possível conectar ao Redis: %v", err))
	}
	fmt.Println("Redis conectado.")
	return rdb
}
