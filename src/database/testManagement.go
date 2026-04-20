package database

import (
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func CriarEntradaProva(prova models.Provas) (uint32, error) {
	var id uint32
	query := "INSERT INTO provas (nome_prova,turma_prova,materia_prova,data_prova) VALUES ($1,$2,$3,$4) RETURNING id;"
	err := DB.QueryRow(query, prova.NomeProva, prova.TurmaProva, prova.MateriaProva, prova.DataProva).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", models.ErroEntradaPostgres, err)
	}
	return id, nil
}

func LerTodasProvas() ([]models.Provas, error) {
	var provas []models.Provas
	rows, err := DB.Query("SELECT id,nome_prova,turma_prova,materia_prova,data_prova FROM provas")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroBuscaPostgres, err)
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Provas
		err := rows.Scan(&p.ID, &p.NomeProva, &p.TurmaProva, &p.MateriaProva, &p.DataProva)
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

func ProcurarProvaNome(nome string) ([]models.Provas, error) {
	var provas []models.Provas
	query := "SELECT id,nome_prova,turma_prova,materia_prova,data_prova FROM provas WHERE nome_prova ILIKE $1"
	rows, err := DB.Query(query, "%"+nome+"%")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroBuscaPostgres, err)
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Provas
		err := rows.Scan(&p.ID, &p.NomeProva, &p.TurmaProva, &p.MateriaProva, &p.DataProva)
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

func UpdateProvas(id uint32, dados models.Provas) error {
	query := `UPDATE provas 
	SET 
    nome_prova = COALESCE(NULLIF($1, ''), nome_prova),
    turma_prova = COALESCE(NULLIF($2, ''), turma_prova),
    materia_prova = COALESCE(NULLIF($3, ''), materia_prova),
    data_prova = COALESCE($4, data_prova)
	WHERE id = $5;`
	if _, err := DB.Exec(query, dados.NomeProva, dados.TurmaProva, dados.MateriaProva, dados.DataProva, id); err != nil {
		return fmt.Errorf("%w: %v", models.ErroAtualizacaoPostgres, err)
	}
	return nil
}

func DeleteProvas(id uint32) error {
	query := `DELETE FROM provas WHERE id=$1`
	res, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroDeletePostgres, err)
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return fmt.Errorf("%w: %d", models.ErroDeleteNenhumaProvaPostgres, id)
	}
	return nil
}
