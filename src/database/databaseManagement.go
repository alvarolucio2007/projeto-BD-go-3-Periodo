// Package database gerencia a abertura e operações com a Base de Dados
package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed sql/init.sql
var initSQL string
var DB *sql.DB // O ponteiro para a DB em si

func ConectarPostgres() (*sql.DB, error) { // Esta função Conecta e checa a saúde do Postgres imediatamente após conectar.
	var err error
	const dsn string = "host=localhost user=user password=password dbname=pg_projeto_db sslmode=disable" // Tenta se conectar à DB com essas credenciais.
	DB, err = sql.Open("pgx", dsn)                                                                       // O sql.Open apenas valida os argumentos, não abre uma conexão.
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
	_, err := DB.Exec(initSQL)
	if err != nil {
		return fmt.Errorf("%w: %v", models.ErroMigracaoPostgres, err)
	}
	log.Println("Migração realizada com sucesso.")
	return nil
}
