// Package models modela tanto erros quanto entradas no Postgres/gPRC e afins, para que seja mais manutenível.
package models

import "log"

type GerenciadorErros struct {
	Mensagem string
}

func (e GerenciadorErros) Log(err error) {
	log.Printf("⚠️ %s | Detalhes: %v", e.Mensagem, err)
}

func (e GerenciadorErros) Fatal(err error) {
	log.Fatalf("❌ FATAL: %s | Detalhes: %v", e.Mensagem, err)
}

// Erros PostgreSQL
var (
	// Erros Setup e migração
	ErroAberturaPostgres = GerenciadorErros{"Erro ao tentar abrir o Postgres: "}
	ErroConexaoPostgres  = GerenciadorErros{"Erro ao tentar conectar ao Postgres: "}
	ErroMigracaoPostgres = GerenciadorErros{"Erro ao tentar migrar o Postgres: "}

	// Erros de Criação
	ErroEntradaPostgres = GerenciadorErros{"Erro ao criar a entrada: "}

	// Erros de Deleção
	ErroDeletePostgres = GerenciadorErros{"Erro ao deletar item: "}
)
