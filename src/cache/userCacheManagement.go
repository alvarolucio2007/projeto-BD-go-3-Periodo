package cache

import (
	"encoding/json"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func AdicionarUsuarioRedis(codigo uint32, usuario models.Usuario) error {
	codigoStr := fmt.Sprintf("user:%s", codigo)

	jsonData, err := json.Marshal(usuario)
	if err != nil {
		return err
	}
	return RedisClient.Set(Ctx, codigoStr, jsonData, 0).Err()
}

func LerUsuarioRedis(codigo uint32) (models.Usuario, error) {
	codigoStr := fmt.Sprintf("user:%s", codigo)
	return RedisClient.Get(Ctx, codigoStr).Result()
}
