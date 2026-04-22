package database

import (
	"database/sql"
	"errors"
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

func AutenticarUsuario(username string, senha string) (bool, string, string, error) {
	var senhaRecebida string
	var role string
	query := "SELECT password ,role FROM usuarios WHERE username ILIKE $1"
	err := DB.QueryRow(query, username).Scan(&senhaRecebida, &role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, "Nem conta tem! Vai criar!", "", nil
		}
		return false, "", "", fmt.Errorf("%w: %v", models.ErroLoginUsuario, err)
	}
	if senha != senhaRecebida {
		msg := fmt.Sprintf("Não foi possível se conectar pois a senha certa é %s", senhaRecebida)
		return false, msg, "", nil
	}
	if role != "admin" {
		msg := "Acesso negado, apenas admins são aceitos."
		return false, msg, "", nil
	}
	return true, "Acesso concluído", role, nil
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
	query := "SELECT id, username,password, role FROM usuarios WHERE username ILIKE $1;"
	rows, err := DB.Query(query, "%"+nome+"%")
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
	if len(usuarios) == 0 {
		return nil, fmt.Errorf("%w", errors.New("nenhum usuário encontrado"))
	}
	return usuarios, nil
}

func UpdateUsuarios(id uint32, dados models.Usuario) error {
	query := `UPDATE usuarios 
        SET 
            username = COALESCE(NULLIF($1, ''), username),
            password = COALESCE(NULLIF($2, ''), password),
            role     = COALESCE(NULLIF($3, ''), role)
        WHERE id = $4`
	res, err := DB.Exec(query, dados.Username, dados.Password, dados.Role, id)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroAtualizacaoPostgres, err)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return models.ErroAtualizacaoPostgres
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
