package cache

import (
	"encoding/json"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarTestRedis(codigo uint32, test models.Provas) error {
	codigoStr := fmt.Sprintf("test:%s", codigo)

	jsonData, err := json.Marshal(test)
	if err != nil {
		return err
	}
	return RedisClient.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func LerTestRedis(codigo uint32) (models.Provas, error) {
	codigoStr := fmt.Sprintf("test:%s", codigo)
	return RedisClient.Get(Ctx, codigoStr).Result()
}
