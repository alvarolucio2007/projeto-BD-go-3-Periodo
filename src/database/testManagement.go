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

func LerTodasProvas() ([]models.Provas, error) {
	var provas []models.Provas
	rows, err := DB.Query("SELECT id,nome_prova,turma_prova,materia_prova FROM provas")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroBuscaPostgres, err)
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Provas
		err := rows.Scan(&p.ID, &p.NomeProva, &p.TurmaProva, &p.MateriaProva)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroBuscaEscanearPostgres, err)
		}
		provas = append(provas, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return provas, nil
}

func UpdateProvas(id int32, dados models.Provas) error {
	query := `UPDATE provas SET nome_prova=$1,turma_prova=$2,materia_prova=$3 WHERE id=$4`
	if _, err := DB.Exec(query, dados.NomeProva, dados.TurmaProva, dados.MateriaProva, id); err != nil {
		return fmt.Errorf("%w: %v", models.ErroAtualizacaoPostgres, err)
	}
	return nil
}
