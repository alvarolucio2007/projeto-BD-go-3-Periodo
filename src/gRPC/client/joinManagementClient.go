package grpcclient

import (
	"context"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *HubConexoes) doInnerJoin() ([]*models.InnerJoinType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.InnerJoin.InnerJoin(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	result := make([]*models.InnerJoinType, 0, len(res.Itens))
	for _, i := range res.Itens {
		result = append(result, &models.InnerJoinType{Username: i.Username, NomeProva: i.NomeProva, NotaProva: i.NotaProva, DataAplicacao: i.DataProva.AsTime()})
	}
	return result, nil
}

func (h *HubConexoes) doLeftJoin() ([]*models.LeftJoinType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := h.LeftJoin.LeftJoin(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	result := make([]*models.LeftJoinType, 0, len(res.Resultado))
	for _, l := range res.Resultado {
		result = append(result, &models.LeftJoinType{Username: l.Username, NomeProva: l.NomeProva, NotaProva: l.NotaProva, DataAplicacao: l.DataProva.AsTime()})
	}
	return result, nil
}
