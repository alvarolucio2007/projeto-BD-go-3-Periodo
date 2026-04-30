package cache

import (
	"encoding/json"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarNotaRedis(codigo uint32, nota models.Notas) error {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	jsonData, err := json.Marshal(nota)
	if err != nil {
		return err
	}
	return RedisClient.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func LerNotaRedis(codigo uint32) (models.Notas, error) {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return RedisClient.Get(Ctx, codigoStr).Result()
}
