package testpostgres

import (
	"testing"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func TestPostgresProva(t *testing.T) {
	t.Run("Criar prova", func(t *testing.T) {
		TestCriarProva(t)
	})
	t.Run("Ler todas as provas", func(t *testing.T) {
		TestLerTodasProvas(t)
	})
	t.Run("Procurar prova específica", func(t *testing.T) {
		TestProcurarProvaNome(t)
	})
	t.Run("Editar prova", func(t *testing.T) {
		// Teste atualização...
	})
	t.Run("Deletar prova", func(t *testing.T) {
		// Teste deleção...
	})
}

var listaProvas = []models.Provas{
	{ID: 1, NomeProva: "Primeiro", TurmaProva: "Primeira turma", MateriaProva: "Primeira matéria", DataProva: time.Date(2026, 4, 17, 12, 0, 0, 0, time.Local)},
	{ID: 2, NomeProva: "Segundo", TurmaProva: "Segunda turma", MateriaProva: "Segunda Matéria", DataProva: time.Date(2026, 4, 18, 12, 0, 0, 0, time.Local)},
	{ID: 3, NomeProva: "Terceiro", TurmaProva: "Terceira turma", MateriaProva: "Terceira Matéria", DataProva: time.Date(2026, 4, 19, 12, 0, 0, 0, time.Local)},
}

func TestCriarProva(t *testing.T) {
	database.ConectarPostgres()
	database.DB.Exec("DROP TABLE IF EXISTS provas CASCADE")
	database.MigrarPostgres()
	t.Run("Criar prova", func(t *testing.T) {
		for i, p := range listaProvas {
			id, err := database.CriarEntradaProva(p)
			if err != nil {
				t.Errorf("Erro ao adicionar prova: %v", err)
			}
			if id != uint32(i+1) {
				t.Errorf("ID é diferente do esperado! Esperava %d, recebi %d", i+1, id)
			}
		}
	})
}

func TestLerTodasProvas(t *testing.T) {
	t.Run("Ler todos as provas", func(t *testing.T) {
		for i, p := range listaProvas {
			listaRecebida, err := database.LerTodasProvas()
			if err != nil {
				t.Errorf("Erro ao receber dados: %v", err)
			}
			if listaRecebida[i] != p {
				t.Errorf("Os dados não batem! Recebi %v e esperava %v", listaRecebida[i], p)
			}
		}
	})
}

func TestProcurarProvaNome(t *testing.T) {
	t.Run("Procurar uma prova", func(t *testing.T) {
		for _, p := range listaProvas {
			listaRecebida, err := database.ProcurarProvaNome(p.NomeProva)
			if err != nil {
				t.Errorf("Erro ao receber dados: %v", err)
			}
			if listaRecebida[0] != p {
				t.Errorf("Os dados não batem! Recebi %v e esperava %v", listaRecebida[0], p)
			}
		}
	})
}
