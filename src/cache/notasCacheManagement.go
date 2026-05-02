package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarNotaRedis(Ctx context.Context, rdb *redis.Client, codigo uint32, nota *models.Notas) error {
	codigoStr := fmt.Sprintf("nota:%d", codigo)
	jsonData, err := json.Marshal(nota)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func AdicionarTodasNotasRedis(Ctx context.Context, rdb *redis.Client, notas []*models.Provas) error {
	const cacheKey = "notas:all"
	jsonData, err := json.Marshal(&notas)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, cacheKey, jsonData, 10*time.Minute).Err()
}

func LerNotaRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) (*models.Notas, error) {
	codigoStr := fmt.Sprintf("nota:%d", codigo)
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		return nil, err
	}
	var resFatorado *models.Notas
	err = json.Unmarshal([]byte(res), resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}

func DeletarNotaRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) error {
	codigoStr := fmt.Sprintf("nota:%d", codigo)
	return rdb.Del(Ctx, codigoStr).Err()
}
