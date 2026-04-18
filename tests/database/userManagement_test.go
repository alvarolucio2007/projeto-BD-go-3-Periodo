package testpostgres

import (
	"testing"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/database"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

func TestPostgresUsuario(t *testing.T) {
	t.Run("Criar usuário", func(t *testing.T) {
		TestCriarUsuario(t)
	})
	t.Run("Autenticar usuário", func(t *testing.T) {
		TestAutenticarUsuario(t)
	})
	t.Run("Ler todos os usuários", func(t *testing.T) {
		TestLerTodosUsuarios(t)
	})
	t.Run("Ler usuário específico", func(t *testing.T) {
		TestProcurarUsuario(t)
	})
	t.Run("Atualizar usuário específico", func(t *testing.T) {
		TestAtualizarUsuario(t)
	})
	t.Run("Deletar usuário", func(t *testing.T) {
		TestDeletarUsuario(t)
	})
}

var listaUsuario = []models.Usuario{
	{ID: 1, Username: "Primeiro", Password: "Senha", Role: "admin"},
	{ID: 2, Username: "Segundo", Password: "SenhaDois", Role: "aluno"},
	{ID: 3, Username: "Terceiro", Password: "SenhaTres", Role: "professor"},
}

func TestCriarUsuario(t *testing.T) {
	database.ConectarPostgres()
	database.DB.Exec("DROP TABLE IF EXISTS usuarios CASCADE")
	database.DB.Exec("DROP TYPE IF EXISTS role_usuario CASCADE")
	database.MigrarPostgres()
	t.Run("Criar usuário", func(t *testing.T) {
		for i, u := range listaUsuario {
			id, err := database.CriarEntradaUsuario(u)
			if err != nil {
				t.Errorf("Erro ao adicionar usuário: %v", err)
			}
			if id != uint32(i+1) {
				t.Errorf("ID é diferente do esperado! Esperava %d, recebi %d", i+1, id)
			}
		}
	})
}

func TestAutenticarUsuario(t *testing.T) {
	t.Run("Autenticar usuário", func(t *testing.T) {
		status, _, err := database.AutenticarUsuario(listaUsuario[0].Username, listaUsuario[0].Password)
		if err != nil {
			t.Errorf("Erro ao autenticar o usuário, %v", err)
		}
		if !status {
			t.Errorf("Erro: Status errado. Esperado: %v, Recebido: %v", true, status)
		}
	})
}

func TestLerTodosUsuarios(t *testing.T) {
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

func TestProcurarUsuario(t *testing.T) {
	t.Run("Procurar um usuário", func(t *testing.T) {
		for _, u := range listaUsuario {
			listaRecebida, err := database.ProcurarUsuario(u.Username)
			if err != nil {
				t.Errorf("Erro ao receber dados: %v", err)
			}
			if listaRecebida[0] != u {
				t.Errorf("Os dados não batem! Recebi %v e esperava %v", listaRecebida[0], u)
			}
		}
	})
}

func TestAtualizarUsuario(t *testing.T) {
	t.Run("Atualizar um usuário", func(t *testing.T) {
		modeloTeste := models.Usuario{ID: 1, Username: "Teste", Password: "12345", Role: "aluno"}
		if err := database.UpdateUsuarios(listaUsuario[0].ID, modeloTeste); err != nil {
			t.Errorf("Erro ao editar usuário: %v", err)
		}
		usuarios, _ := database.ProcurarUsuario(modeloTeste.Username)
		if usuarios[0] != modeloTeste {
			t.Errorf("Usuário não foi editado corretamente! Esperado: %v, Recebido: %v", modeloTeste, usuarios[0])
		}
	})
}

func TestDeletarUsuario(t *testing.T) {
	t.Run("Deletar um usuário", func(t *testing.T) {
		for _, u := range listaUsuario {
			if err := database.DeleteUsuarios(u.ID); err != nil {
				t.Errorf("Erro ao deletar o usuário: %v", err)
			}
		}
		resultado, err := database.LerTodosUsuarios()
		if err != nil {
			t.Errorf("Erro ao pegar a lista vazia dos usuários. %v", err)
		}
		if len(resultado) != 0 {
			t.Errorf("Ainda há resultados na DB: achados %v", resultado)
		}
	})
}
