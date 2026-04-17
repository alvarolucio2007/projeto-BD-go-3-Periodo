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
	{Username: "Primeiro", Password: "Senha", Role: "Admin"},
	{Username: "Segundo", Password: "SenhaDois", Role: "Estudante"},
	{Username: "Terceiro", Password: "SenhaTres", Role: "Professor"},
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
}
