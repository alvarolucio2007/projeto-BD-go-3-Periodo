package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

var DB *sql.DB

func ConectarDatabasePostgres() (*sql.DB, error) {
	var err error
	const dsn string = "host=pg_projeto_db user=user password=password dbname=shortener sslmode=disable"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroAberturaPostgres, err)
	}
	for i := range 5 {
		if err = DB.Ping(); err == nil {
			return DB, nil
		}
		log.Printf("Aguardando Postgres: (Tentativa %d/5)", i+1)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("%w: %v", models.ErroConexaoPostgres, err)
}
