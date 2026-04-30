package cache

import (
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarNotaRedis(codigo uint32, nota models.Notas) error {
	codigoStr := strconv.FormatUint(uint64(codigo), 32)
	return RedisClient.Set(Ctx, codigoStr, nota, 0).Err()
}

func LerNotaRedis(codigo uint32) (models.Usuario, error) {
	codigoStr := strconv.FormatUint(uint64(codigo), 32)
	return RedisClient.Get(Ctx, codigoStr).Result
}
