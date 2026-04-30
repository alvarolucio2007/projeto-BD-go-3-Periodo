package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarTestRedis(Ctx context.Context, rdb *redis.Client, codigo uint32, test models.Provas) error {
	codigoStr := fmt.Sprintf("test:%s", codigo)
	jsonData, err := json.Marshal(test)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func LerTestRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) (models.Provas, error) {
	codigoStr := fmt.Sprintf("test:%s", codigo)
	return rdb.Get(Ctx, codigoStr).Result()
}
