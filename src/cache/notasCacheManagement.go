package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarNotaRedis(Ctx context.Context, rdb *redis.Client, codigo uint32, nota *models.Notas) error {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	jsonData, err := json.Marshal(nota)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func LerNotaRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) (*models.Notas, error) {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	res := rdb.Get(Ctx, codigoStr).Return()
	var resFatorado *models.Notas
	err := json.Unmarshal(res, resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}

func DeletarUsuarioRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) error {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return rdb.Del(Ctx, codigoStr).Err()
}
