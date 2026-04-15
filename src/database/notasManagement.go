package database

import (
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func CriarEntradaNotas(user models.Notas) (int32, error) {
	var id int32
	query := "INSERT INTO notas (usuario_id,prova_id,nota_prova) VALUES ($1,$2,$3) RETURNING id;"
	if err := DB.QueryRow(query, user.UsuarioID, user.ProvaID, user.NotaProva).Scan(&id); err != nil {
		return 0, fmt.Errorf("%w: %v", models.ErroEntradaPostgres, err)
	}
	return id, nil
}
