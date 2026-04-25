// Package database gerencia a abertura e operações com a Base de Dados
package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed sql/initNotas.sql
var initSQLNotas string

//go:embed sql/initProvas.sql
var initSQLProvas string

//go:embed sql/initUsuarios.sql
var initSQLUsuarios string
var DB *sql.DB // O ponteiro para a DB em si

func ConectarPostgres() (*sql.DB, error) { // Esta função Conecta e checa a saúde do Postgres imediatamente após conectar.
	var err error
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		dsn = "postgres://user_escola:password_escola@localhost:5432/db_escola?sslmode=disable"
	}
	DB, err = sql.Open("pgx", dsn) // O sql.Open apenas valida os argumentos, não abre uma conexão.
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroAberturaPostgres, err)
	}
	for i := range 5 { // Checa 5 vezes em 10 segundos (1 vez a cada 2 segundos) para que se certifique que a DB não está viva.
		if err = DB.Ping(); err == nil { // Pinga a DB, para checar a vida dela, se não retorna erro, retorna o ponteiro da DB
			return DB, nil
		}
		log.Printf("Aguardando Postgres: (Tentativa %d/5)", i+1) // Se não, tenta novamente em 2s
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("%w: %v", models.ErroConexaoPostgres, err) // Após 10s, o sistema decide que a DB está realmente morta e retorna o erro.
}

func MigrarPostgres() error {
	_, err := DB.Exec(initSQLUsuarios)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroMigracaoPostgres, err)
	}
	_, err = DB.Exec(initSQLProvas)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroMigracaoPostgres, err)
	}
	_, err = DB.Exec(initSQLNotas)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroMigracaoPostgres, err)
	}
	_, err = DB.Exec("INSERT INTO usuarios (username, password, role) VALUES ('root','admin','admin')")
	if err != nil {
		fmt.Errorf("Erro de migração: %v ", err)
	}
	log.Println("Migração realizada com sucesso.")
	return nil
}
