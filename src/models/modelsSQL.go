package models

type Usuario struct {
	ID       int32
	Username string
	Password string
	Role     string
}
type Provas struct {
	ID           int32
	NomeProva    string
	TurmaProva   string
	MateriaProva string
}
type Notas struct {
	ID        int32
	UsuarioID int32
	ProvaID   int32
	NotaProva float32
}
