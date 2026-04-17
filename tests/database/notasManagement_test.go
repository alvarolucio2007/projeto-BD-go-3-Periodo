package testpostgres

import (
	"fmt"
	"testing"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func TestPostgresNotas(t *testing.T) {
	t.Run("Criar nota", func(t *testing.T) {
		TestCriarNotas(t)
	})
	t.Run("Buscar notas", func(t *testing.T) {
		TestBuscarNotas(t)
	})
	t.Run("Atualizar nota", func(t *testing.T) {
		//
	})
	t.Run("Deletar nota", func(t *testing.T) {
		//
	})
}

var listaNotas = []models.Notas{
	{ID: 1, UsuarioID: 1, ProvaID: 1, NotaProva: 1},
	{ID: 2, UsuarioID: 2, ProvaID: 2, NotaProva: 2},
	{ID: 3, UsuarioID: 3, ProvaID: 3, NotaProva: 3},
}

func TestCriarNotas(t *testing.T) {
	database.ConectarPostgres()
	_, err := database.DB.Exec("DROP TABLE IF EXISTS notas CASCADE")
	if err != nil {
		t.Fatalf("Erro ao limpar banco para testes: %v", err)
	}
	_, err = database.DB.Exec("DROP TABLE IF EXISTS provas CASCADE")
	if err != nil {
		t.Fatalf("Erro ao limpar banco para testes: %v", err)
	}
	_, err = database.DB.Exec("DROP TABLE IF EXISTS usuarios CASCADE")
	if err != nil {
		t.Fatalf("Erro ao limpar banco para testes: %v", err)
	}
	database.MigrarPostgres()

	for i := range listaNotas {
		// Cria um usuário genérico para o teste (ID vai ser gerado automaticamente pelo RESTART IDENTITY)
		_, err := database.CriarEntradaUsuario(models.Usuario{
			Username: fmt.Sprintf("UserTest%d", i+1),
			Password: "123",
			Role:     "aluno",
		})
		if err != nil {
			t.Fatalf("Falha ao criar usuário de dependência: %v", err)
		}

		// Cria uma prova genérica para o teste
		_, err = database.CriarEntradaProva(models.Provas{
			NomeProva:    "Prova Teste",
			MateriaProva: "Go",
		})
		if err != nil {
			t.Fatalf("Falha ao criar prova de dependência: %v", err)
		}
	}
	t.Run("Criar nota", func(t *testing.T) {
		for i, n := range listaNotas {
			id, err := database.CriarEntradaNotas(n)
			if err != nil {
				t.Errorf("Erro ao adicionar a nota: %v", err)
			}
			if id != uint32(i+1) {
				t.Errorf("ID é diferente do esperado! Esperava %d, recebi %d", i+1, id)
			}

		}
	})
}

func TestBuscarNotas(t *testing.T) {
	database.DB.Exec("TRUNCATE TABLE notas, provas, usuarios RESTART IDENTITY CASCADE")
	t.Run("Buscar notas geral", func(t *testing.T) {
		for i, n := range listaNotas {
			database.CriarEntradaUsuario(models.Usuario{Username: fmt.Sprintf("User%d", i), Password: "1234", Role: "aluno"})
			database.CriarEntradaProva(models.Provas{NomeProva: fmt.Sprintf("Prova%d", i), TurmaProva: "1", MateriaProva: "1"})
			database.CriarEntradaNotas(n)
		}
		listaRecebida, _ := database.BuscarNotas("")
		if len(listaRecebida) != len(listaNotas) {
			t.Errorf("Os dados não batem! Recebi %d e esperava %d", len(listaRecebida), len(listaNotas))
		}
	})
	t.Run("Buscar nota específica", func(t *testing.T) {
		usuario := models.Usuario{Username: "BuscaSpec", Role: "aluno"}
		prova := models.Provas{NomeProva: "ProvaSpec"}
		uID, _ := database.CriarEntradaUsuario(usuario)
		pID, _ := database.CriarEntradaProva(prova)
		notaValor := float32(10.0)
		database.CriarEntradaNotas(models.Notas{UsuarioID: uID, ProvaID: pID, NotaProva: notaValor})
		result, _ := database.BuscarNotas(usuario.Username)
		if len(result) == 0 {
			t.Fatalf("Vazio.")
		}
		resultado := result[0]
		if resultado.NotaProva != notaValor {
			t.Errorf("Nota incorreta! Esperava %v, recebi %v", notaValor, resultado.NotaProva)
		}
	})
}
