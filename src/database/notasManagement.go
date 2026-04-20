package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func CriarEntradaNotas(nota models.Notas) (uint32, error) {
	var id uint32
	query := "INSERT INTO notas (usuario_id,prova_id,nota_prova) SELECT ($1,$2,$3) FROM usuarios WHERE id=$1 AND role='aluno' RETURNING id;"
	if err := DB.QueryRow(query, nota.UsuarioID, nota.ProvaID, nota.NotaProva).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("usuário %d não é aluno ou não existe", nota.UsuarioID)
		}
		return 0, fmt.Errorf("%w: %v", models.ErroEntradaPostgres, err)
	}
	return id, nil
}

func BuscarNotas(username string) ([]models.InnerJoinType, error) {
	var notas []models.InnerJoinType

	// SQL Base
	query := `SELECT u.username, p.nome_prova, n.nota_prova, p.data_prova
        FROM notas n
        INNER JOIN usuarios u ON u.id = n.usuario_id
        INNER JOIN provas p ON p.id = n.prova_id`

	var rows *sql.Rows
	var err error

	if username != "" {
		query += ` WHERE u.username ILIKE $1`
		rows, err = DB.Query(query, "%"+username+"%")
	} else {
		rows, err = DB.Query(query)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroBuscaPostgres, err)
	}
	defer rows.Close()

	for rows.Next() {
		var n models.InnerJoinType
		// O Scan agora é sempre o mesmo!
		if err := rows.Scan(&n.Username, &n.NomeProva, &n.NotaProva, &n.DataAplicacao); err != nil {
			return nil, fmt.Errorf("%w, %v", models.ErroBuscaEscanearPostgres, err)
		}
		notas = append(notas, n)
	}
	return notas, nil
}

func UpdateNotas(id uint32, dados models.Notas) error {
	query := `UPDATE notas SET usuario_id=$1,prova_id=$2,nota_prova=$3 WHERE id=$4;`
	if _, err := DB.Exec(query, dados.UsuarioID, dados.ProvaID, dados.NotaProva, id); err != nil {
		return fmt.Errorf("%w: %v", models.ErroAtualizacaoPostgres, err)
	}
	return nil
}

func DeleteNotas(id uint32) error {
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
