package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarTestRedis(Ctx context.Context, rdb *redis.Client, codigo uint32, test *models.Provas) error {
	codigoStr := fmt.Sprintf("test:%d", codigo)
	jsonData, err := json.Marshal(&test)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func AdicionarTodosTestRedis(Ctx context.Context, rdb *redis.Client, tests []*models.Provas) error {
	const cacheKey = "test:all"
	jsonData, err := json.Marshal(&tests)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, cacheKey, jsonData, 10*time.Minute).Err()
}

func LerTestRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) (*models.Provas, error) {
	codigoStr := fmt.Sprintf("test:%d", codigo)
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		return nil, err
	}
	var resFatorado models.Provas
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return &resFatorado, nil
}

func LerAllTestRedis(Ctx context.Context, rdb *redis.Client) ([]*models.Provas, error) {
	const cacheKey = "test:all"
	res, err := rdb.Get(Ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}
	var resFatorado []*models.Provas
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}

func DeletarTestRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) error {
	codigoStr := fmt.Sprintf("user:%d", codigo)
	return rdb.Del(Ctx, codigoStr).Err()
}
