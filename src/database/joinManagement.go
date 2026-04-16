package database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

//go:embed sql/innerJoin.sql
var innerJoinSQL string

//go:embed sql/leftJoin.sql
var leftJoinSQL string

type LeftJoinType struct {
	Username      string
	NomeProva     sql.NullString
	NotaProva     sql.NullFloat64
	DataAplicacao time.Time
}

func LeftJoin() ([]LeftJoinType, error) {
	rows, err := DB.Query(leftJoinSQL)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinExecutarPostgres, err)
	}
	defer rows.Close()
	var listaLeftJoins []LeftJoinType
	for rows.Next() {
		var r LeftJoinType
		if err := rows.Scan(&r.Username, &r.NomeProva, &r.NotaProva, &r.DataAplicacao); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinScanPostgres, err)
		}
		listaLeftJoins = append(listaLeftJoins, r)
	}
	return listaLeftJoins, nil
}

type InnerJoinType struct {
	Username      string
	NomeProva     string
	NotaProva     float32
	DataAplicacao time.Time
}

func InnerJoin() ([]InnerJoinType, error) {
	rows, err := DB.Query(innerJoinSQL)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroInnerJoinExecutarPostgres, err)
	}
	defer rows.Close()
	var listaInnerJoins []InnerJoinType
	for rows.Next() {
		var i InnerJoinType
		if err := rows.Scan(&i.Username, &i.NomeProva, &i.NotaProva, &i.DataAplicacao); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroInnerJoinScanPosgres, err)
		}
		listaInnerJoins = append(listaInnerJoins, i)
	}
	return listaInnerJoins, nil
}
