package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarUsuarioRedis(Ctx context.Context, rdb *redis.Client, codigo uint32, usuario *models.Usuario) error {
	codigoStr := fmt.Sprintf("user:%d", codigo)
	jsonData, err := json.Marshal(&usuario)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 10*time.Minute).Err()
}

func LerUsuarioRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) (*models.Usuario, error) {
	codigoStr := fmt.Sprintf("user:%d", codigo)
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		return nil, err
	}
	var resFatorado models.Usuario
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return &resFatorado, nil
}

func AdicionarTodosUsuariosRedis(Ctx context.Context, rdb *redis.Client, users []*models.Usuario) error {
	const cacheKey = "user:all"
	jsonData, err := json.Marshal(&users)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, cacheKey, jsonData, 10*time.Minute).Err()
}

func LerTodosUsuariosRedis(Ctx context.Context, rdb *redis.Client) ([]*models.Usuario, error) {
	const cacheKey = "user:all"
	val, err := rdb.Get(Ctx, cacheKey).Result()
	if err != nil {
		return nil, err
	}
	var usuarioLista []*models.Usuario
	if err := json.Unmarshal([]byte(val), &usuarioLista); err != nil {
		return nil, err
	}
	return usuarioLista, nil
}

func DeletarUsuarioRedis(Ctx context.Context, rdb *redis.Client, codigo uint32) error {
	codigoStr := fmt.Sprintf("user:%d", codigo)
	return rdb.Del(Ctx, codigoStr).Err()
}
