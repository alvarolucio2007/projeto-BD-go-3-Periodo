package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

func AdicionarLerQuantidadeProvaAlunoTodos(Ctx context.Context, rdb *redis.Client, data map[string]int64) error {
	codigoStr := "quantidade_prova_aluno"
	jsonData, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	return rdb.Set(Ctx, codigoStr, jsonData, 10*time.Minute).Err()
}
