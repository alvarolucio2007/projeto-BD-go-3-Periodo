// Package models modela tanto erros quanto entradas no Postgres/gPRC e afins, para que seja mais manutenível.
package models

import "errors"

// Erros PostgreSQL
var (
	// Erros Setup e migração
	ErroAberturaPostgres = errors.New("erro ao tentar abrir o Postgres: ")
	ErroConexaoPostgres  = errors.New("erro ao tentar conectar ao Postgres: ")
	ErroMigracaoPostgres = errors.New("erro ao tentar migrar o Postgres: ")

	// Erros de Criação
	ErroEntradaPostgres = errors.New("erro ao criar a entrada: ")

	// Erros de Deleção
	ErroDeletePostgres = errors.New("erro ao deletar item: ")

	// Erros de Procura
	ErroBuscaPostgresNEncontrado = errors.New("não foram encontrados itens: ")
)
