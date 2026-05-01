package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarUsuarioRedis(Ctx context.Context, rdb *redis.Client, codigo uint32, usuario *models.Usuario) error {
	codigoStr := fmt.Sprintf("user:%s", codigo)
	jsonData, err := json.Marshal(&usuario)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func LerUsuarioRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) (*models.Usuario, error) {
	codigoStr := fmt.Sprintf("user:%s", codigo)
	res := RedisClient.Get(Ctx, codigoStr).Return()
	var resFatorado *models.Usuario
	err := json.Unmarshal(res, resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}

func DeletarUsuarioRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) error {
	codigoStr := fmt.Sprintf("user:%s", codigo)
	return rdb.Del(Ctx, codigoStr).Err()
}
