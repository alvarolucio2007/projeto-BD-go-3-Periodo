// Package models modela tanto erros quanto entradas no Postgres/gPRC e afins, para que seja mais manutenível.
package models

import "errors"

// Erros PostgreSQL
var (
	// Erros Setup e migração
	ErroAberturaPostgres       = errors.New("erro ao tentar abrir o Postgres")
	ErroConexaoPostgres        = errors.New("erro ao tentar conectar ao Postgres")
	ErroLeituraArquivoMigracao = errors.New("erro ao tentar ler o arquivo de migração do Postgres")
	ErroMigracaoPostgres       = errors.New("erro ao tentar migrar o Postgres")
	// Erro Login usuário
	ErroLoginUsuario = errors.New("erro ao tentar fazer login")
	// Erro busca de senha KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK
	ErroBuscaSenha           = errors.New("erro ao buscar a senha no postgres")
	ErroUsuarioNaoEncontrado = errors.New("usuário não encontrado")

	// Erros de Criação (C)
	ErroEntradaPostgres = errors.New("erro ao tentar criar a entrada")

	// Erros de Procura (R)
	ErroBuscaPostgres         = errors.New("erro ao tentar buscar")
	ErroBuscaEscanearPostgres = errors.New("erro ao tentar escanear o modelo")

	// Erros de Atualização (U)
	ErroAtualizacaoPostgres = errors.New("erro ao tentar atualizar o item")

	// Erros de Deleção (D)
	ErroDeletePostgres              = errors.New("erro ao deletar item")
	ErroDeleteNenhumUsuarioPostgres = errors.New("nenhum usuário encontrado com o ID ")
	ErroDeleteNenhumaProvaPostgres  = errors.New("nenhuma prova encontrada com o ID")

	// Erros de Left Join
	ErroLeftJoinLerArquivoPostgres = errors.New("erro ao ler arquivo de left join")
	ErroLeftJoinExecutarPostgres   = errors.New("erro ao executar o comando de left join")
	ErroLeftJoinScanPostgres       = errors.New("erro ao dar scan no resultado do left join")

	// Erros de Inner Join
	ErroInnerJoinLerArquivoPostgres = errors.New("erro ao ler arquivo de inner join")
	ErroInnerJoinExecutarPostgres   = errors.New("erro ao executar o comando de inner join")
	ErroInnerJoinScanPosgres        = errors.New("erro ao dar scan no resultado do inner join")
)
