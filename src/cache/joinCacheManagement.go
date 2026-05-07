package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarInnerJoinRedis(Ctx context.Context, rdb *redis.Client, innerJoins []*models.InnerJoinType) error {
	const codigoStr = "innerJoin:all"
	jsonData, err := json.Marshal(&innerJoins)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 10*time.Minute).Err()
}

func LerInnerJoinRedis(Ctx context.Context, rdb *redis.Client) ([]*models.InnerJoinType, error) {
	const codigoStr = "innerJoin:all"
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		return nil, err
	}
	var resFatorado []*models.InnerJoinType
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}

func AdicionarLeftJoinRedis(Ctx context.Context, rdb *redis.Client, leftJoins []*models.LeftJoinType) error {
	const codigoStr = "leftJoin:all"
	jsonData, err := json.Marshal(&leftJoins)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 10*time.Minute).Err()
}

func LerLeftJoinRedis(Ctx context.Context, rdb *redis.Client) ([]*models.LeftJoinType, error) {
	const codigoStr = "leftJoin:all"
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		return nil, err
	}
	var resFatorado []*models.LeftJoinType
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}
