package cache

import "github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"

func AdicionarUsuarioRedis(codigo uint32, usuario models.Usuario) error {
	return RedisClient.Set(Ctx, codigo, usuario, 0).Err()
}

func LerUsuarioRedis(codigo uint32) (models.Usuario, error) {
	return RedisClient.Get(Ctx, codigo).Result()
}
