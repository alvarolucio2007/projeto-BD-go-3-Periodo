package cache

import (
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarNotaRedis(codigo uint32, nota models.Notas) error {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return RedisClient.Set(Ctx, codigoStr, nota, 0).Err()
}

func LerNotaRedis(codigo uint32) (models.Usuario, error) {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return RedisClient.Get(Ctx, codigoStr).Result
}
