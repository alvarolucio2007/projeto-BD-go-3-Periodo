package testpostgres

import (
	"testing"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func TestPostgresProva(t *testing.T) {
	t.Run("Criar prova", func(t *testing.T) {
		// Teste criação...
	})
	t.Run("Ler todas as provas", func(t *testing.T) {
		// Teste leitura geral...
	})
	t.Run("Procurar prova específica", func(t *testing.T) {
		// Teste procura...
	})
	t.Run("Editar prova", func(t *testing.T) {
		// Teste atualização...
	})
	t.Run("Deletar prova", func(t *testing.T) {
		// Teste deleção...
	})
}

var listaProvas = []models.Provas{
	{ID: 1, NomeProva: "Primeiro", TurmaProva: "Primeira turma", MateriaProva: "Primeira matéria", DataProva: time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)},
	{ID: 2, NomeProva: "Segundo", TurmaProva: "Segunda turma", MateriaProva: "Segunda Matéria", DataProva: time.Date(2026, 4, 18, 12, 0, 0, 0, time.UTC)},
	{ID: 3, NomeProva: "Terceiro", TurmaProva: "Terceira turma", MateriaProva: "Terceira Matéria", DataProva: time.Date(2026, 4, 19, 12, 0, 0, 0, time.UTC)},
}

func TestCriarProva(t *testing.T) {
	database.ConectarPostgres()
	database.DB.Exec("DROP TABLE IF EXISTS provas CASCADE")
}
