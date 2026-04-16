package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

type LeftJoinType struct {
	Username  string
	NomeProva sql.NullString
	NotaProva sql.NullFloat64
}

func LeftJoin() ([]LeftJoinType, error) {
	queryByte, err := os.ReadFile("src/sql/leftJoin.sql")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinLerArquivoPostgres, err)
	}
	queryString := string(queryByte)
	rows, err := DB.Query(queryString)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinExecutarPostgres, err)
	}
	defer rows.Close()
	var listaLeftJoins []LeftJoinType
	for rows.Next() {
		var r LeftJoinType
		if err := rows.Scan(&r.Username, &r.NomeProva, &r.NotaProva); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinScanPostgres, err)
		}
		listaLeftJoins = append(listaLeftJoins, r)
	}
	return listaLeftJoins, nil
}

type InnerJoinType struct {
	Username  string
	NomeProva string
	NotaProva float32
}

func InnerJoin() ([]InnerJoinType, error) {
	queryByte, err := os.ReadFile("src/sql/fullJoin.sql")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroInnerJoinLerArquivoPostgres, err)
	}
	queryString := string(queryByte)
	rows, err := DB.Query(queryString)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroInnerJoinExecutarPostgres, err)
	}
	defer rows.Close()
	var listaInnerJoins []InnerJoinType
	for rows.Next() {
		var i InnerJoinType
		if err := rows.Scan(&i.Username, &i.NomeProva, &i.NotaProva); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroInnerJoinScanPosgres, err)
		}
		listaInnerJoins = append(listaInnerJoins, i)
	}
	return listaInnerJoins, nil
}
