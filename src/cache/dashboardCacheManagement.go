package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

func AdicionarQuantidadeProvaAlunos(Ctx context.Context, rdb *redis.Client, data map[string]int64) error {
	codigoStr := "quantidade_prova_aluno"
	jsonData, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 10*time.Minute).Err()
}

func LerQuantidadeProvaAlunos(Ctx context.Context, rdb *redis.Client) (map[string]int64, error) {
	codigoStr := "quantiade_prova_aluno"
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	var resFatorado map[string]int64
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}

func AdicionarQuantidadeNotaProvaAlunos(Ctx context.Context, rdb *redis.Client, dados map[string]models.EstatisticaAluno) error {
	codigoStr := "quantidade_nota_prova_aluno"
	jsonData, err := json.Marshal(&dados)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 10*time.Minute).Err()
}

func LerQuantidadeNotaProvaAlunos(Ctx context.Context, rdb *redis.Client) (map[string]models.EstatisticaAluno, error) {
	codigoStr := "quantidade_nota_prova_aluno"
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	var resFatorado map[string]models.EstatisticaAluno
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}

func AdicionarMediaNotaMaterias(Ctx context.Context, rdb *redis.Client, dados map[string]models.EstatisticaAluno) error {
	codigoStr := "medias_notas"
	jsonData, err := json.Marshal(&dados)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 10*time.Minute).Err()
}

func LerMediaNotaMaterias(Ctx context.Context, rdb *redis.Client) (map[string]models.EstatisticaAluno, error) {
	codigoStr := "medias_notas"
	res, err := rdb.Get(Ctx, codigoStr).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	var resFatorado map[string]models.EstatisticaAluno
	err = json.Unmarshal([]byte(res), &resFatorado)
	if err != nil {
		return nil, err
	}
	return resFatorado, nil
}
