package database

import (
	_ "embed"
	"fmt"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

//go:embed sql/innerJoin.sql
var innerJoinSQL string

//go:embed sql/leftJoin.sql
var leftJoinSQL string

func LeftJoin() ([]models.LeftJoinType, error) {
	rows, err := DB.Query(leftJoinSQL)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinExecutarPostgres, err)
	}
	defer rows.Close()
	var listaLeftJoins []models.LeftJoinType
	for rows.Next() {
		var l models.LeftJoinType
		if err := rows.Scan(&l.Username, &l.NomeProva, &l.NotaProva, &l.DataAplicacao); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinScanPostgres, err)
		}
		listaLeftJoins = append(listaLeftJoins, l)
	}
	return listaLeftJoins, nil
}

func InnerJoin() ([]models.InnerJoinType, error) {
	rows, err := DB.Query(innerJoinSQL)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErroInnerJoinExecutarPostgres, err)
	}
	defer rows.Close()
	var listaInnerJoins []models.InnerJoinType
	for rows.Next() {
		var i models.InnerJoinType
		if err := rows.Scan(&i.Username, &i.NomeProva, &i.NotaProva, &i.DataAplicacao); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroInnerJoinScanPosgres, err)
		}
		listaInnerJoins = append(listaInnerJoins, i)
	}
	return listaInnerJoins, nil
}
