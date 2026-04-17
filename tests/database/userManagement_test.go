package testpostgres

import (
	"testing"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func TestPostgresUsuario(t *testing.T) {
	t.Run("Criar entrada", func(t *testing.T) {
		// teste de entrada...
	})
	t.Run("Ler todos os usuários", func(t *testing.T) {
		// teste de leitura geral...
	})
	t.Run("Ler usuário específico", func(t *testing.T) {
		// teste de leitura específico...
	})
	t.Run("Atualizar usuário específico", func(t *testing.T) {
		// teste de atualizar usuários...
	})
	t.Run("Deletar usuário", func(t *testing.T) {
		// Teste de deletar usuários...
	})
}

var listaUsuario = []models.Usuario{
	{ID: 1, Username: "Primeiro", Password: "Senha", Role: "Admin"},
	{ID: 2, Username: "Segundo", Password: "SenhaDois", Role: "Estudante"},
	{ID: 3, Username: "Terceiro", Password: "SenhaTres", Role: "Professor"},
}

func TestCriarUsuario(t *testing.T) {
	t.Run("Criar usuário", func(t *testing.T) {
		for i, u := range listaUsuario {
			id, err := database.CriarEntradaUsuario(u)
			if err != nil {
				t.Errorf("Erro ao adicionar usuário: %v", err)
			}
			if id != uint32(i+1) {
				t.Errorf("ID é diferente do esperado! Esperava %d, recebi %d", id, i+1)
			}
		}
	})
	t.Run("Ler todos os usuários", func(t *testing.T) {
		for i, u := range listaUsuario {
			listaRecebida, err := database.LerTodosUsuarios()
			if err != nil {
				t.Errorf("Erro ao receber dados: %v", err)
			}
			if listaRecebida[i] != u {
				t.Errorf("Os dados não batem! Recebi %v e esperava %v", listaRecebida[i], u)
			}
		}
	})
}
