package cache

import (
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarTestRedis(codigo uint32, test models.Provas) {
	codigoStr := strconv.FormatUint(uint64(codigo), 32)
	return RedisClient.Set(Ctx, codigoStr, test, 0).Err()
}

func LerTestRedis(codigo uint32) (models.Provas, error) {
	codigoStr := strconv.FormatUint(uint64(codigo), 32)
	return RedisClient.Get(Ctx, codigoStr).Result()
}
