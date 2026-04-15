package database

import (
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func CriarEntradaProva(prova models.Provas) (int32, error) {
	var id int32
	query := "INSERT INTO provas (nome_prova,turma_prova,materia_prova) VALUES ($1,$2,$3) RETURNING id;"
	err := DB.QueryRow(query, prova.NomeProva, prova.TurmaProva, prova.MateriaProva).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", models.ErroEntradaPostgres, err)
	}
	return id, nil
}
