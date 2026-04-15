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
	ErroAberturaPostgres = GerenciadorErros{"Erro ao tentar abrir o Postgres: "}
	ErroConexaoPostgres  = GerenciadorErros{"Erro ao tentar conectar ao Postgres: "}
)
