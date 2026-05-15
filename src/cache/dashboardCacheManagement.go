package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/redis/go-redis/v9"
)

const (
	KeyQuantidadeProvaAluno     = "dashboard:quantidade_prova_aluno"
	KeyQuantidadeNotaProvaAluno = "dashboard:quantidade_nota_prova_aluno"
	KeyMediasNotas              = "dashboard:medias_notas"
	KeyDistribuicaoStatus       = "dashboard:distribuicao_status"
	DefaultTTL                  = 10 * time.Minute
)

func AdicionarQuantidadeProvaAlunos(Ctx context.Context, rdb *redis.Client, data map[string]int64) error {
	jsonData, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, KeyQuantidadeProvaAluno, jsonData, 10*time.Minute).Err()
}

func LerQuantidadeProvaAlunos(Ctx context.Context, rdb *redis.Client) (map[string]int64, error) {
	res, err := rdb.Get(Ctx, KeyQuantidadeProvaAluno).Result()
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
	jsonData, err := json.Marshal(&dados)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, KeyQuantidadeNotaProvaAluno, jsonData, 10*time.Minute).Err()
}

func LerQuantidadeNotaProvaAlunos(Ctx context.Context, rdb *redis.Client) (map[string]models.EstatisticaAluno, error) {
	res, err := rdb.Get(Ctx, KeyQuantidadeNotaProvaAluno).Result()
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
	jsonData, err := json.Marshal(&dados)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, KeyMediasNotas, jsonData, 10*time.Minute).Err()
}

func LerMediaNotaMaterias(Ctx context.Context, rdb *redis.Client) (map[string]models.EstatisticaAluno, error) {
	res, err := rdb.Get(Ctx, KeyMediasNotas).Result()
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

func AdicionarDistribuicaoStatusAluno(Ctx context.Context, rdb *redis.Client, dados map[string]int64) error {
	jsonData, err := json.Marshal(&dados)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, KeyDistribuicaoStatus, jsonData, 10*time.Minute).Err()
}

func LerDistribuicaoStatusAluno(Ctx context.Context, rdb *redis.Client) (map[string]int64, error) {
	dados, err := rdb.Get(Ctx, KeyDistribuicaoStatus).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	var result map[string]int64
	err = json.Unmarshal([]byte(dados), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
