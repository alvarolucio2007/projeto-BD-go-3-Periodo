package models

import (
	"time"
)

type Usuario struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Provas struct {
	ID           uint32    `json:"id"`
	NomeProva    string    `json:"nome_prova"`
	TurmaProva   string    `json:"turma_prova"`
	MateriaProva string    `json:"materia_prova"`
	DataProva    time.Time `json:"data_prova"`
}

type Notas struct {
	ID        uint32  `json:"id"`
	UsuarioID uint32  `json:"usuario_id"`
	ProvaID   uint32  `json:"prova_id"`
	NotaProva float32 `json:"nota_prova"`
}

type InnerJoinType struct {
	Username      string    `json:"username"`
	NomeProva     string    `json:"nome_prova"`
	NotaProva     float32   `json:"nota_prova"`
	DataAplicacao time.Time `json:"data_aplicacao"`
}

type LeftJoinType struct {
	Username      string     `json:"username"`
	NomeProva     *string    `json:"nome_prova"`
	NotaProva     *float32   `json:"nota_prova"`
	DataAplicacao *time.Time `json:"data_aplicacao"`
}

type AuthResult struct {
	Status   bool   `json:"status"`
	Mensagem string `json:"mensagem"`
	Role     string `json:"role"`
}
