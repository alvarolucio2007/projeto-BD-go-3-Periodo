package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *InnerJoinConexao) DoInnerJoin() ([]*models.InnerJoinType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := i.InnerJoin.InnerJoin(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	result := make([]*models.InnerJoinType, 0, len(res.Itens))
	for _, i := range res.Itens {
		result = append(result, &models.InnerJoinType{Username: i.Username, NomeProva: i.NomeProva, NotaProva: i.NotaProva, DataAplicacao: i.DataProva.AsTime()})
	}
	return result, nil
}

func (l *LeftJoinConexao) DoLeftJoin() ([]*models.LeftJoinType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := l.LeftJoin.LeftJoin(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	result := make([]*models.LeftJoinType, 0, len(res.Resultado))
	for _, l := range res.Resultado {
		t := l.DataProva.AsTime()
		result = append(result, &models.LeftJoinType{Username: l.Username, NomeProva: l.NomeProva, NotaProva: l.NotaProva, DataAplicacao: &t})
	}
	return result, nil
}
