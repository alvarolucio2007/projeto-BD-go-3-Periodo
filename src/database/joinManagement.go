package database

import (
	"fmt"
	"os"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
)

type LeftJoinType struct {
	username  string
	nomeProva string
	notaProva float32
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
		if err := rows.Scan(&r.username, &r.notaProva, &r.nomeProva); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErroLeftJoinScanPostgres, err)
		}
		listaLeftJoins = append(listaLeftJoins, r)
	}
	return listaLeftJoins, nil
}
