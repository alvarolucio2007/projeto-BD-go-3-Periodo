package models

import (
	"time"
)

type Usuario struct {
	ID       uint32
	Username string
	Password string
	Role     string
}
type Provas struct {
	ID           uint32
	NomeProva    string
	TurmaProva   string
	MateriaProva string
	DataProva    time.Time
}
type Notas struct {
	ID        uint32
	UsuarioID uint32
	ProvaID   uint32
	NotaProva float32
}
type InnerJoinType struct {
	Username      string
	NomeProva     string
	NotaProva     float32
	DataAplicacao time.Time
}
type LeftJoinType struct {
	Username      string
	NomeProva     *string
	NotaProva     *float32
	DataAplicacao *time.Time
}
type AuthResult struct {
	Status   bool
	Mensagem string
	Role     string
}
