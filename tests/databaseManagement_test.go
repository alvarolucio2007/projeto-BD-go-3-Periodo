package maintest

import (
	"database/sql"
	"testing"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
)

func TestPostgres(t *testing.T) {
	t.Run("Conexão e Migração", func(t *testing.T) {
		testConexaoMigracao(t)
	})
}

var db *sql.DB

func testConexaoMigracao(t *testing.T) {
	var err error
	t.Run("Conectar DB", func(t *testing.T) {
		db, err = database.ConectarPostgres()
		if err != nil {
			t.Errorf("Erro ao conectar database: %v", err)
		}
	})
	t.Run("Migrar DB", func(t *testing.T) {
		if err = database.MigrarPostgres(db); err != nil {
			t.Errorf("Erro ao migrar database: %v", err)
		}
	})
}

func testUser(t *testing.T) {
}
