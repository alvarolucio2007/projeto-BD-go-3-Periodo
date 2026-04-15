package database

import (
	"database/sql"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func CriarEntradaUsuario(username string, password string, role string) (int32, error) {
	var id int32
	query := "INSERT INTO usuario (username,password,role) VALUES ($1,$2,$3) RETURNING id;"
	err = DB.QueryRow(query, username, password, role).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", models.ErroEntradaPostgres, err)
	}
	return id, nil
}

func LerTodosUsuarios(db *sql.DB) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroBuscaPostgres, err)
	}
	defer rows.Close()
	for rows.Next() {
		var u models.Usuario
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Role)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroBuscaEscanearPostgres, err)
		}
		usuarios = append(usuarios, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usuarios, nil
}

func UpdateUsuarios(id int32, dados models.Usuario) error {
	query := `UPDATE usuarios SET username=$1, password=$2, role=$3 WHERE id=$4`
	if _, err := DB.Exec(query, dados.Username, dados.Password, dados.Role, id); err != nil {
		return fmt.Errorf("%w: %v", models.ErroAtualizacaoPostgres, err)
	}
	return nil
}
