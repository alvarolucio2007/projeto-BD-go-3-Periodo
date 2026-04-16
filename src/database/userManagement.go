package database

import (
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func CriarEntradaUsuario(user models.Usuario) (uint32, error) {
	var id uint32
	query := "INSERT INTO usuarios (username,password,role) VALUES ($1,$2,$3) RETURNING id;"
	err := DB.QueryRow(query, user.Username, user.Password, user.Role).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", models.ErroEntradaPostgres, err)
	}
	return uint32(id), nil
}

func LerTodosUsuarios() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	rows, err := DB.Query("SELECT id,username,password,role FROM usuarios")
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

func ProcurarUsuario(nome string) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	query := "SELECT id, username,password, role FROM usuarios WHERE username=$1"
	rows, err := DB.Query(query, nome)
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

func UpdateUsuarios(dados models.Usuario) error {
	query := `UPDATE usuarios SET username=$1, password=$2, role=$3 WHERE id=$4`
	if _, err := DB.Exec(query, dados.Username, dados.Password, dados.Role, dados.ID); err != nil {
		return fmt.Errorf("%w: %v", models.ErroAtualizacaoPostgres, err)
	}
	return nil
}

func DeleteUsuarios(id uint32) error {
	query := `DELETE FROM usuarios WHERE id=$1`
	res, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroDeletePostgres, err)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("%w: %d", models.ErroDeleteNenhumUsuarioPostgres, id)
	}
	return nil
}
