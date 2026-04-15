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

func LerTodasNotas() ([]models.Notas, error) {
	var notas []models.Notas
	rows, err := DB.Query("SELECT id,usuario_id,prova_id,nota_prova FROM provas")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroBuscaPostgres, err)
	}
	defer rows.Close()
	for rows.Next() {
		var n models.Notas
		if err := rows.Scan(&n.ID, n.UsuarioID, n.ProvaID, n.NotaProva); err != nil {
			return nil, fmt.Errorf("%w, %v", models.ErroBuscaEscanearPostgres, err)
		}
		notas = append(notas, n)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return notas, nil
}

func UpdateNotas(id int32, dados models.Notas) error {
	query := `UPDATE notas SET usuario_id=$1,prova_id=$2,nota_prova=$3 WHERE id=$4;`
	if _, err := DB.Exec(query, dados.UsuarioID, dados.ProvaID, dados.NotaProva); err != nil {
		return fmt.Errorf("%w: %v", models.ErroAtualizacaoPostgres, err)
	}
	return nil
}

func DeleteNotas(id int32) error {
	query := `DELETE FROM notas where id=$1`
	res, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroDeletePostgres, err)
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return fmt.Errorf("%w: %d", models.ErroDeleteNenhumaProvaPostgres, id)
	}
	return nil
}
