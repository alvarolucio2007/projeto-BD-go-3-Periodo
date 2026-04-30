package cache

import (
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarTestRedis(codigo uint32, test models.Provas) {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return RedisClient.Set(Ctx, codigoStr, test, 0).Err()
}

func LerTestRedis(codigo uint32) (models.Provas, error) {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return RedisClient.Get(Ctx, codigoStr).Result()
}
