package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConectarDatabasePostgres() error {
	var err error
	const dsn string = "host=pg_projeto_db user=user password=password dbname=shortener sslmode=disable"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err
}
