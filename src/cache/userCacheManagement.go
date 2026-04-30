package cache

import (
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarUsuarioRedis(codigo uint32, usuario models.Usuario) error {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return RedisClient.Set(Ctx, codigoStr, usuario, 0).Err()
}

func LerUsuarioRedis(codigo uint32) (models.Usuario, error) {
	codigoStr := fmt.Sprintf("nota:%s", codigo)
	return RedisClient.Get(Ctx, codigoStr).Result()
}
